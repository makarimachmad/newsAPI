<template>
  <v-app>
    <v-app-bar
      app
      color="light-grey"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://spesolution.com/images/layouts/main/logo.png"
          transition="scale-transition"
          width="80"
        />

        News
      </div>

      <v-spacer></v-spacer>

    </v-app-bar>

    <v-main>
      <router-view></router-view>
      <br><br><br>
      <Footer/>
    </v-main>
  </v-app>
</template>

<script>
import Footer from './components/Footer.vue';

export default {
  name: 'App',

  components: {
    Footer
  },

  data: () => ({
    //
  }),
};
</script>