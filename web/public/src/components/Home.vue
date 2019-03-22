<template>
    <div class="row justify-content-center">
        <div class="col-sm-6">
            <h1>PUBG FUN STATS</h1>
            <b-input-group prepend="Username" class="mt-3">
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
                    <th></th>
                </tr>
                <tr v-for="m in matches">
                    <td>
                        <span class="badge badge-secondary">{{ m.GameMode }}</span>
                    </td>
                    <td>
                        #{{ player[m.ID].Stats.WinPlace }}/{{ m.Rosters.length }}<br>
                        <span v-if="player[m.ID].Stats.WinPlace == 1" class="badge badge-warning">Winner!</span>
                    </td>
                    <td>{{ Math.floor(player[m.ID].Stats.TimeSurvived / 60)}} minutes</td>
                    <td>{{ player[m.ID].Stats.Kills }}</td>
                    <td>{{ parseInt(player[m.ID].Stats.DamageDealt) }}</td>
                    <td>{{ new Date(m.CreatedAt).toDateString() }}</td>
                    <td>{{ mapNames[m.MapName] }}</td>
                    <td>
                        <span v-for="p in rosters[m.ID].Participants">{{ p.Stats.Name }}<br></span>
                    </td>
                    <td>
                        <b-button v-b-modal="'modal'+m.ID">More</b-button>
                    </td>
                    <td>
                        <b-button v-b-modal="'modal'+m.ID">Kill Stats</b-button>
                    </td>
                    <b-modal :id="'modal'+m.ID" title="BootstrapVue">
                        <table class="table table-striped table-bordered" v-if="matches.length" v-cloak>
                            <tr>
                                <th>Name</th>
                                <th>Place</th>
                                <th>Time Survived</th>
                                <th>Kill</th>
                                <th>Damage</th>
                            </tr>
                            <tr v-for="p in participants[m.ID]">
                                <td>{{ p.Stats.Name }}</td>
                                <td>{{ p.Stats.WinPlace }}</td>
                                <td>{{ Math.floor(p.Stats.TimeSurvived / 60)}} minutes</td>
                                <td>{{ p.Stats.Kills }}</td>
                                <td>{{ parseInt(p.Stats.DamageDealt) }}</td>
                            </tr>
                        </table>
                    </b-modal>
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
                player: {},
                participants: {},
                telemetry: {},
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
                                        this.player[m.ID] = p;
                                        this.rosters[m.ID] = r;
                                    }
                                    if (!this.participants[m.ID]) {
                                        this.participants[m.ID] = []
                                    }
                                    this.participants[m.ID].push(p)
                                })
                            })
                        })
                    }
                    return this.matches;
                }
            },
            getTelemetry: async function (url) {
                if (url) {
                    let t = (await axios.get('api/telemetry/', {
                        params: {endpointURL: url}
                    })).data.data;
                    this.telemetry[t.MatchID] = t;
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
