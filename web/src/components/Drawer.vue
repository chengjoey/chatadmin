<template>
    <div>
    <!-- <v-sheet
        color="grey lighten-4"
        class="pa-4"
      >
        <v-avatar
        color="teal"
        size="64"
        >
        <span class="white--text headline">{{ currentUser.nickname}}</span>
        </v-avatar>
        <span></span>
        

      </v-sheet>

      <v-divider></v-divider> -->

      <!-- <v-list>
        <v-list-item
          v-for="[icon, text, path] in links"
          :key="icon"
          :to="path"
          :link=true
        >
          <v-list-item-icon>
            <v-icon>{{ icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list> -->
      <v-list>
        <v-list-item
          to="/app/dashboard"
          :link=true
          v-if="role == 'admin'"
        >
        <v-list-item-icon>
            <v-icon>mdi-inbox-arrow-down</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>Dashboard</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item
          to="/app/chat"
          :link=true
        >
        <v-list-item-icon>
            <v-icon>mdi-send</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>Chat</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </div>
</template>

<script>
// import { mapState } from 'vuex'
import jwtService from '@/common/jwt.service'
export default {
    data: () => ({
      links: [
        ['mdi-inbox-arrow-down', 'Dashboard', '/app/dashboard'],
        ['mdi-send', 'Chat', '/app/chat'],
      ],
      role: ""
    }),
    methods: {
        onSubmit(path) {
            console.log(path)
            this.$router.push({"name": path})
        },
        isAdmin() {
          this.role = jwtService.getRole();
        }
    },
    mounted() {
      this.isAdmin();
  },
  }
</script>