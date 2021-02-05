<template>
  <v-card
    class="mx-auto"
  >
    <v-card-title class="title font-weight-regular justify-space-between">
      <span>Login</span>
    </v-card-title>

    <v-window>
      <v-window-item>
        <v-card-text>
          <v-text-field
            label="Nickname"
            v-model="nickname"
          ></v-text-field>
          <span class="caption grey--text text--darken-1">
            admin, adminstrator为管理员用户，其它为普通用户
          </span>
        </v-card-text>
      </v-window-item>
      </v-window>

    <v-divider></v-divider>

    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn
        color="primary"
        depressed
        @click="login"
      >
        login
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
    import { LOGIN } from "@/store/actions.type";
    import { mapState } from "vuex";

  export default {
    data: () => ({
      nickname: ""
    }),
    methods: {
        login () {
        this.$store
            .dispatch(LOGIN, this.nickname)
            .then(() => {
              if (this.user.role == "admin") {
                this.$router.push({ path: "/app/dashboard" })
              } else {
                this.$router.push({ path: "/app/chat" })
              }
            })
      },
    },
    computed: {
        ...mapState({
            user: state => state.auth.user
        })
    }
  }
</script>

<style>
.mx-auto {
    width: 300px;
    height: 300px;
    position: center;
    top: 25%;
}
</style>