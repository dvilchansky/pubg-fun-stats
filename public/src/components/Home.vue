<template>
    <div class="row justify-content-center">
        <div class="col-sm-6">
            <h1>PUBG FUN STATS</h1>
            <b-input-group size="lg" prepend="Username" class="mt-3">
                <b-form-input v-model="username"/>
                <b-input-group-append>
                    <b-button @click="getMatches" variant="outline-success">Button</b-button>
                </b-input-group-append>
            </b-input-group>
            <b-spinner label="Spinning" style="display: none"/>
            <br>
            <table class="table table-striped table-bordered" v-if="matches.length" v-cloak>
                <tr>
                    <th>Game Mode</th>
                    <th>Duration</th>
                    <th>Date</th>
                    <th>Map Name</th>
                    <th></th>
                </tr>
                <tr v-for="m in matches" v-bind:key="m.match_id">
                    <td><span class="badge badge-secondary">{{ m.game_mode }}</span></td>
                    <td>{{ Math.floor(m.duration / 60) }} minutes</td>
                    <td>{{ new Date(m.created_at).toDateString() }}</td>
                    <td>{{ mapNames[m.map_name] }}</td>
                    <td>
                        <b-button variant="outline-success">more info</b-button>
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
                matches: [],
                username: ''
            }
        },
        name: 'Home',
        methods: {
            getMatches: async function () {
                if (this.username) {

                    return this.matches = (await axios.post('api/players/' + this.username)).data.data;
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
