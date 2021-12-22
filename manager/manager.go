package manager

import (
	"context"
	"fmt"
	"github.com/nate-droid/go-orchestra/messages"
	"golang.org/x/sync/errgroup"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/nate-droid/go-orchestra/core/symphony"
	"time"
)

type Manager struct {
	WaitForSymphony chan *symphony.Symphony
	SymphonyReady   chan bool
}

func NewManager() (*Manager, error) {
	ec, err := messages.NewEncodedNatsCon()
	if err != nil {
		return nil, err
	}

	recvCh := make(chan *symphony.Symphony)
	_, err = ec.BindRecvChan("createSymphony", recvCh)
	if err != nil {
		return nil, err
	}
	sendCh := make(chan bool)
	err = ec.BindSendChan("symphonyReady", sendCh)
	if err != nil {
		return nil, err
	}

	man := &Manager{
		WaitForSymphony: recvCh,
		SymphonyReady:   sendCh,
	}

	return man, nil
}

func (m *Manager) hireOrchestra(symphony *symphony.Symphony) {
	// for each section, we need to hire the musicians
	for _, section := range symphony.Sections {
		for i := 0; i < section.GroupSize; i++ {
			// create Job / Container
			fmt.Println("Hired!")
		}
	}
	// TODO send orchestra ready
	// sleep for 2 seconds then send
	time.Sleep(3 * time.Second)
	m.SymphonyReady <- true
}

func (m *Manager) hireMusician(song *symphony.SongStructure) bool {
	// create a container here!
	return true
}

func (m *Manager) Run(ctx context.Context) error {
	clientset, err := connectToCluster()
	if err != nil {
		return err
	}
	//err = createJob(clientset)
	//
	//if err != nil {
	//	return err
	//}
	fmt.Println("Manager entering loop")
	errs, ctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		for {
			fmt.Println("waiting for Symphony")
			select {
			case s := <-m.WaitForSymphony:
				err = createJob(clientset, s)
				if err != nil {
					return err
				}
				fmt.Println("completed job")
			}

		}
	})

	return errs.Wait()
}

func connectToCluster() (*kubernetes.Clientset, error) {
	// TODO this can be deleted if the below works
	//home, exists := os.LookupEnv("HOME")
	//if !exists {
	//	home = "/root"
	//}
	//
	//// TODO create service-account and config
	//
	//configPath := filepath.Join(home, ".kube", "config")
	//config, err := clientcmd.BuildConfigFromFlags("", configPath)
	//if err != nil {
	//	return nil, err
	//}
	fmt.Println("Connecting to cluster")
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	fmt.Println("creating config")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func createJob(clientset *kubernetes.Clientset, symphony *symphony.Symphony) error {
	// TODO make these nicer
	imageName := "natedroid/go-orchestra:latest"
	// imageCount := 1
	symphonyID := "symphony-1"
	serviceType := "musician"
	namespace := v1.NamespaceDefault
	jobName := symphony.SymphonyID

	job := clientset.BatchV1().Jobs(namespace)
	var backOffLimit int32 = 0

	//containers := make([]v1.Container, imageCount)
	//
	//for i := 0; i < imageCount; i++ {
	//	c := v1.Container{
	//		Name:  fmt.Sprintf("%s-musician-%d", symphonyID, i),
	//		Image: imageName,
	//		Env: []v1.EnvVar{
	//			{
	//				Name:  "SERVICE_TYPE",
	//				Value: serviceType,
	//			},
	//		},
	//	}
	//	containers = append(containers, c)
	//}

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
			Labels: map[string]string{
				"app":          "go-orchestra",
				"service-type": serviceType,
			},
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &backOffLimit,
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: jobName,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            fmt.Sprintf("%s-musician-%d", symphonyID, 3),
							Image:           imageName,
							ImagePullPolicy: v1.PullNever,
							Env: []v1.EnvVar{
								{
									Name:  "SERVICE_TYPE",
									Value: serviceType,
								},
								{
									Name:  "NATS_URI",
									Value: "nats://nats:4222",
								},
								{
									Name:  "SINGLE_RUN",
									Value: "true",
								},
							},
						},
					},
					RestartPolicy:      v1.RestartPolicyOnFailure,
					ServiceAccountName: "default",
				},
			},
		},
	}

	_, err := job.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Println("Job created")

	return nil
}
