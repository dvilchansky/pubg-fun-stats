<template>
    <div>
        <h1>{{ msg }}</h1>
        <input type="text" v-model="username">
        <button @click="getMatches">submit</button>
        <table v-if="matches" border="5px">
            <tr v-for="m in matches" v-bind:key="m.match_id">
                <td>{{ m.match_id }}</td>
                <td>{{ m.game_mode }}</td>
                <td>{{ m.duration }}</td>
                <td>{{ m.created_at }}</td>
                <td>{{ m.map_name }}</td>
            </tr>
        </table>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        data() {
            return {
                matches: [],
                username: ''
            }
        },
        name: 'HelloWorld',
        props: {
            msg: String
        },
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
