<template>
    <div class="row justify-content-center">
        <div class="col-sm-6">
            <h1>PUBG FUN STATS</h1>
            <b-input-group size="lg" prepend="Username" class="mt-3">
                <b-form-input v-model="username"/>
                <b-input-group-append>
                    <b-button @click="getMatches" variant="outline-success">Search</b-button>
                </b-input-group-append>
            </b-input-group>
            <b-spinner label="Spinning" style="display: none"/>
            <br>
            <table class="table table-striped table-bordered" v-if="matches.length" v-cloak>
                <tr>
                    <th>Game Mode</th>
                    <th>Place</th>
                    <th>Time Survived</th>
                    <th>Kill</th>
                    <th>Damage</th>
                    <th>Date</th>
                    <th>Map</th>
                    <th>Roster</th>
                    <th></th>
                </tr>
                <tr v-for="m in matches" v-bind:key="m.ID">
                    <td>
                        <span class="badge badge-secondary">{{ m.GameMode }}</span>
                    </td>
                    <td>
                        #{{ participants[m.ID].Stats.WinPlace }}/{{ m.Rosters.length }}<br>
                        <span v-if="participants[m.ID].Stats.WinPlace == 1" class="badge badge-warning">Winner!</span>
                    </td>
                    <td>{{ Math.floor(participants[m.ID].Stats.TimeSurvived / 60)}} minutes</td>
                    <td>{{ participants[m.ID].Stats.Kills }}</td>
                    <td>{{ parseInt(participants[m.ID].Stats.DamageDealt) }}</td>
                    <td>{{ new Date(m.CreatedAt).toDateString() }}</td>
                    <td>{{ mapNames[m.MapName] }}</td>
                    <td>
                        <span v-for="p in rosters[m.ID].Participants" v-bind:key="p.ID">{{ p.Stats.Name }}<br></span>
                    </td>
                    <td>
                        <b-button variant="outline-success">more</b-button>
                    </td>
                </tr>
            </table>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        data() {
            return {
                mapNames: {
                    "Desert_Main": "Miramar",
                    "DihorOtok_Main": "Vikendi",
                    "Erangel_Main": "Erangel",
                    "Range_Main": "Camp Jackal",
                    "Savage_Main": "Sanhok"
                },
                rosters: {},
                participants: {},
                matches: [],
                username: ''
            }
        },
        name: 'Home',
        methods: {
            getMatches: async function () {
                if (this.username) {
                    this.matches = (await axios.get('api/players/' + this.username)).data.data;
                    if (this.matches.length) {
                        this.matches.forEach((m) => {
                            m.Rosters.forEach((r) => {
                                r.Participants.forEach((p) => {
                                    if (p.Stats.Name === this.username) {
                                        this.participants[m.ID] = p;
                                        this.rosters[m.ID] = r;
                                    }
                                })
                            })
                        })
                    }
                    return this.matches;
                }
            }
        }
    }


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h3 {
        margin: 40px 0 0;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }
</style>
