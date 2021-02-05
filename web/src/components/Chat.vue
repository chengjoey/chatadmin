<template>
  <v-container class="fill-height">
    <section ref="chatArea" class="chat-area">
      <p
        v-for="(msg, index) in historyMsgList"
        class="message"
        :class="{
          'message-out': msg.user.nickname === currentUser.nickname,
          'message-in': msg.user.nickname !== currentUser.nickname,
        }"
        :key="`hsmsg-${index}`"
      >
        {{ msg.user.nickname }} at {{formatDate(msg.msg_time)}}
        <br /><span>{{msg.content}}</span>
      </p>
      <br />
      <p align="center">以上为历史消息</p>
      <v-divider></v-divider>
      <br />
      <p
        v-for="(msg, index) in msglist"
        class="message"
        :class="{
          'message-out': msg.user.nickname === currentUser.nickname,
          'message-in': msg.user.nickname !== currentUser.nickname,
        }"
        :key="`msg-${index}`"
      >
        {{ msg.user.nickname }} at {{formatDate(msg.msg_time)}}
        <br /><span>{{msg.content}}</span>
      </p>
    </section>
    <v-col cols="12">
      <v-text-field
        v-model="content"
        append-outer-icon="mdi-send"
        filled
        clear-icon="mdi-close-circle"
        clearable
        label="Message"
        type="text"
        @click:append="toggleMarker"
        @click:append-outer="sendMessage"
        @click:clear="clearMessage"
      ></v-text-field>
    </v-col>
  </v-container>
</template>
<script>
import { mapGetters } from "vuex";
import ApiService from "@/common/api.service";
import { WS_URL } from "@/common/config";

export default {
  data: () => ({
    websock: null,
    msglist: [],
    historyMsgList: [],
    content: "Hey!",
    client: "",

    show: false,
    message: "Hey!",
    marker: true,
  }),
  methods: {
    joinChat: function () {
      let token = encodeURIComponent(ApiService.getIDToken());
      let url = WS_URL + "?token=" + token + "&client=" + this.client;
      // let decodeUrl = decodeURIComponent(url)
      console.log(url);
      this.websock = new WebSocket(url);
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onopen = this.websocketonopen;
      this.websock.onerror = this.websocketonerror;
      this.websock.onclose = this.websocketonclose;
    },
    websocketonmessage(e) {
      const data = JSON.parse(e.data);
      data.receive_time = new Date();
      this.msglist.push(data);
    },
    websocketonopen() {
      console.log("websocket已建立");
    },
    websocketonerror() {
      console.log("websocket断开链接");
    },
    websocketonsend() {
      console.log(this.currentUser.nickname)
      let msg = JSON.stringify({ content: this.content });
      this.websock.send(msg);
      let data = {
        user: {
          nickname: this.currentUser.nickname,
          uid: this.currentUser.uid,
        },
        type: 0,
        content: this.content,
        msg_time: new Date().getTime(),
      };

      this.msglist.push(data);
      this.content = "";
    },
    websocketonclose(e) {
      console.log("断开链接", e);
    },
    closewebsocket() {
      this.websock.close();
    },
    // 换行
    lineFeed: function (evt) {
      console.log(evt);
      this.content = this.content + "\n";
    },
    formatDate: function (dateStr) {
      let d = new Date(dateStr);
      return d.toLocaleString();
    },
    toggleMarker() {
      this.marker = !this.marker;
    },
    sendMessage() {
      this.websocketonsend();
      this.clearMessage();
    },
    clearMessage() {
      this.message = "";
    },
    resetIcon() {
      this.iconIndex = 0;
    },
    changeIcon() {
      this.iconIndex === this.icons.length - 1
        ? (this.iconIndex = 0)
        : this.iconIndex++;
    },
    getMsgs() {
      const path = "v1/msgs?client=" + this.client;
      ApiService.get(path)
        .then((response) => {
          this.historyMsgList = response.data.value;
        })
        .catch((response) => {
          console.log(response)
        })
    },
    getParams() {
      this.client = this.$route.params.client
    },
    printUser() {
      console.log(this.currentUser);
    },
  },
  mounted() {
    this.printUser();
    this.joinChat();
  },
  computed: {
    ...mapGetters(["currentUser"])
  },
  destroyed() {
    this.websock.close();
  },
  created() {
    this.getParams()
    this.getMsgs();
  },
  beforeRouteLeave(to, from, next){
    this.websock.close()
    next()
  }
};
</script>

<style scoped>
.chat-area {
  border: 1px solid #ccc;
  background: rgb(255, 255, 255);
  height: 60vh;
  padding: 1em;
  width: 150vh;
  overflow: auto;
  /* max-width: 1000px; */
  margin: 0 auto 2em auto;
  box-shadow: 2px 2px 5px 2px rgba(0, 0, 0, 0.3);
}
.message {
  width: 45%;
  border-radius: 10px;
  padding: 0.5em;
  font-size: 0.8em;
}
/* .message-out {
  background: #407fff;
  color: white;
  margin-left: 50%;
}
.message-in {
  background: #f1f0f0;
  color: black;
} */
.message-out {
  background: #407fff;
  color: white;
  margin-left: 50%;
}
.message-in {
  background: #f1f0f0;
  color: black;
}
</style>