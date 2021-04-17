<template>
  <q-layout view="lHh Lpr lFf">
    <q-tabs
        v-model="tab"
        align="left"
        narrow-indicator
    >
      <q-tab class="text-purple" name="mails" label="Mails" />
      <q-tab class="text-orange" name="alarms" label="Alarms" />
      <q-tab class="text-teal" name="movies" label="Movies" />
    </q-tabs>

    <q-tab-panels
        animated
        v-model="tab"
        transition-prev="scale"
        transition-next="scale"
        class="text-primary text-left"
    >
      <q-tab-panel name="mails">
        <Main />
      </q-tab-panel>

      <q-tab-panel name="alarms">
        <div class="text-h6">Alarms</div>
        Ad molestiae non facere animi nobis, similique nemo velit reiciendis corporis impedit nam in.
      </q-tab-panel>

      <q-tab-panel name="movies">
        <div class="text-h6">Movies</div>
        Nostrum necessitatibus expedita dolores? Voluptatem repudiandae magni ea.
      </q-tab-panel>
    </q-tab-panels>
  </q-layout>
</template>

<script>
import Main from "@/components/Main";
// import Info from "@/components/Info";
// import Setting from "@/components/Setting";
// import About from "@/components/About";

export default {
  name: 'app',

  components: {
    // Main, Info, Setting, About
    Main
  },

  data() {
    return {
      tab: 'mails'
    }
  }
}
</script>

<style>
</style>
