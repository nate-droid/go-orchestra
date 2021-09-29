<script>
    import { onMount } from "svelte";
    export let results;
    onMount(async () => {
        let x = [];
        fetch('http://127.0.0.1:1323/symphonies').then(res => res.json()).then(data => {
            data.results.forEach(song => {
                try {
                    let chords = [];
                    let intervals = [];
                    if (song.ChordProgression !== null && song.Progression !== null){
                        song.ChordProgression.forEach(c => {
                            if (c !== null) {
                                chords.push(c);
                            }
                        })
                        song.Progression.forEach(p => {
                            if (p !== null) {
                                intervals.push(p);
                            }
                        })
                        song.ChordProgression = chords;
                        song.Progression = intervals;
                        x.push(song);
                    }
                }
                catch (e) {
                    console.log("oops: ", e);
                }

            });
            results = x;
            loading = false;
        })
    });
    export let loading = true;
</script>

<div class="bg-white px-4 py-5 border-b border-gray-200 sm:px-6">
    <h3 class="text-lg leading-6 font-medium text-gray-900">
        Recent Songs
    </h3>
</div>

<div class="flex flex-col">
    <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
            <div class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
                <table class="min-w-full divide-y divide-gray-200">
                    <thead class="bg-gray-50">
                    <tr>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            SymphonyID
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Chords
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Intervals
                        </th>
                        <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                            Band Section
                        </th>
                        <th scope="col" class="relative px-6 py-3">
                            <span class="sr-only">Edit</span>
                        </th>
                    </tr>
                    </thead>
                    <tbody>
                    {#if loading}
                        loading...
                    {:else}
                        {#each results as song, index }
                            <tr class=" {index % 2  ? 'bg-white' : 'bg-gray-50'}">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                                    {song.SymphonyID}
                                </td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                                    {#if song.ChordProgression !== null}
                                        {#each song.ChordProgression as prog}

                                            {prog.Root}{prog.Name}  {"  "}
                                        {/each}
                                    {/if}
                                </td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                                    {#each song.Progression as interval}
                                        {interval} {" "}
                                    {/each}
                                </td>
                                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                                    {song.Type}
                                </td>
                            </tr>
                        {/each}
                    {/if}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>