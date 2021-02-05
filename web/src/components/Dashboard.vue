<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <strong>在线用户</strong>
      </v-col>
    </v-row>
    <v-row>
      <template v-for="n in line">
        <!-- <v-col
          :key="n"
          class="mt-2"
          cols="12"
        >
          <strong>Category {{ n }}</strong>
        </v-col> -->

        <v-col
          v-for="(user, j) in users.slice((n-1)*4, n*4)"
          :key="`${n}${j}`"
          cols="6"
          md="2"
        >
          <v-card class="mx-auto" max-width="344" outlined>
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-title class="headline mb-1">
                  {{ user.nickname }}
                </v-list-item-title>
                <v-list-item-subtitle v-if="user.is_chated"
                  >已在聊天</v-list-item-subtitle
                >
                <v-list-item-subtitle v-else>待聊天</v-list-item-subtitle>
                <v-list-item-subtitle>加入时间: {{ user.enter_at }}</v-list-item-subtitle>
              </v-list-item-content>

              <!-- <v-list-item-avatar
                tile
                size="80"
                color="grey"
              ></v-list-item-avatar> -->
            </v-list-item>

            <v-card-actions>
              <v-btn outlined rounded text :disabled="user.is_chated" @click="chat(user.nickname)"> 点击聊天 </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </template>
    </v-row>
    <!-- <v-row>
      <template>
        <v-col
          v-for="(user, index) in users"
          :key="`${index}`"
        >
          <v-card class="mx-auto" max-width="344" outlined>
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-title class="headline mb-1">
                  {{ user.nickname }}
                </v-list-item-title>
                <v-list-item-subtitle v-if="user.is_chated"
                  >已在聊天</v-list-item-subtitle
                >
                <v-list-item-subtitle v-else>待聊天</v-list-item-subtitle>
              </v-list-item-content>

              <v-list-item-avatar
                tile
                size="80"
                color="grey"
              ></v-list-item-avatar>
            </v-list-item>

            <v-card-actions>
              <v-btn outlined rounded text> Button </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </template>
    </v-row> -->
  </v-container>
</template>

<script>
import { API_URL } from "@/common/config";
import ApiService from "@/common/api.service";
import EventSourcePolify from "eventsource";

export default {
  data: () => ({
    users: [],
    line: 0,
    evtSource: null,
  }),
  methods: {
    createSSE() {
      var EventSource = EventSourcePolify;
      this.evtSource = new EventSource(API_URL + "v1/stream", {
        headers: { Authorization: ApiService.getIDToken() },
      });
      this.evtSource.onmessage = this.sseOnMessage;
    },
    getUsers() {
      ApiService.get("v1/users")
        .then((response) => {
          this.users = response.data.value;
          this.line = Math.ceil(this.users.length / 4);
        })
        .catch((response) => {
          console.log(response)
        })
    },
    chat(nickname) {
      this.$router.push({
        paht: '/app/chat',
        name: 'chat',
        params: {
          client: nickname
        }
      })
    },
    sseOnMessage(e) {
      if (e.data === "null") {
        this.users = [];
        this.line = 0;
      } else {
        this.users = JSON.parse(e.data);
        this.line = Math.ceil(this.users.length / 4);
      }
      console.log(this.users, this.line)
    },
  },
  mounted() {
    this.getUsers();
    this.createSSE();
  },
};
</script>