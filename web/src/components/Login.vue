<template>
  <validation-observer
    ref="observer"
    v-slot="{ invalid }"
  >
    <form @submit.prevent="submit">
      <validation-provider
        v-slot="{ errors }"
        name="Name"
        rules="required|max:10"
      >
        <v-text-field
          v-model="name"
          :counter="10"
          :error-messages="errors"
          label="Name"
          required
        ></v-text-field>
      </validation-provider>

      <v-btn
        class="mr-4"
        type="submit"
        :disabled="invalid"
      >
        submit
      </v-btn>
      <v-btn @click="clear">
        clear
      </v-btn>
    </form>
  </validation-observer>
</template>

<script>
  import { mapState } from "vuex";
  import { LOGIN } from "@/store/actions.type";
  export default {
    data: () => ({
      name: '',
        }),
    methods: {
      submit () {
        this.$store
            .dispatch(LOGIN, this.name)
            .then(() => {this.$router.push({ name: "dashboard" })})
      },
      clear () {
        this.name = ''
      },
    },
    computed: {
        ...mapState({
            errors: state => state.auth.errors
        })
    }
  }
</script>

