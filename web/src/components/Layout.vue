<template>
    <div class="app-wrapper">
        <v-system-bar app>
        <v-spacer></v-spacer>
<!-- 
        <v-icon>mdi-square</v-icon>

        <v-icon>mdi-circle</v-icon>

        <v-icon>mdi-triangle</v-icon> -->
        </v-system-bar>

        <v-app-bar app color="white" flat>
            <v-app-bar-nav-icon @click="drawer = !drawer">
            </v-app-bar-nav-icon>
            <v-toolbar-title>ChatAdmin</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn v-on:click="Logout()" dark>
                <v-icon
                dark
                left
                >
                mdi-minus-circle
                </v-icon>Logout
            </v-btn>
        </v-app-bar>

        <v-navigation-drawer
        v-model="drawer"
        app
        >
        <Drawer/> 
        </v-navigation-drawer>

        <v-main>
        <v-container
            class="py-8 px-6"
            fluid
        >
        <Appmain/>
        </v-container>
        </v-main>
    </div>
</template>

<script>
  import Drawer from '@/components/Drawer'
  import Appmain from '@/components/Appmain'
    import { mapGetters } from 'vuex'
  import { LOGOUT } from "@/store/actions.type"

  export default {
    data: () => ({
      drawer: null,
      setRight: true,
    }),
    components: {
      Drawer,
      Appmain
    },
    methods: {
        Logout: function() {
            this.$store
                .dispatch(LOGOUT)
                .then(() => this.$router.push({ path: "/login" }))
        }
    },
    computed: {
      ...mapGetters(["currentUser"])
    }
  }
</script>

<style>
.bar-container {
    width: 300px;
    color: darkblue;
}
</style>