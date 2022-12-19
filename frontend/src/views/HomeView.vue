<script lang="ts">
import Card from "@/components/Card.vue";
export default {
  name: "home",
  components: { Card },
  methods: {
    redirect(path: string) {
      this.$router.push(path);
    },
  },
};
</script>

<template>
  <main class="h-screen grid grid-cols-4 items-center gap-10 content-center">
    <Card
      img="https://img.freepik.com/premium-vector/hight-electric-power-transmissian-tower-line-high-volt-isometric-nature-cartoon_255805-174.jpg?w=2000"
      icon="fa-solid fa-bolt"
      detail="Relation of Electricity and Temperature"
      name="Electic-Temp"
      @click="redirect('/elect-temp')"
    />
    <Card
      class="w-full"
      img="https://img.freepik.com/premium-vector/water-drop-fire-mascot-cartoon-vector_193274-19977.jpg?w=2000"
      icon="fa-solid fa-plug"
      detail="Relation of Water and Temperature"
      name="Water-Temp"
      @click="redirect('/water-temp')"
    />
    <Card
      img="https://img.freepik.com/premium-vector/water-splash-with-ripple-effect-cartoon-blue-droplets-isolated-white-background_81894-6084.jpg?w=2000"
      icon="fa-solid fa-water"
      detail="Relation of Water and Rainfall"
      name="Water-Rain"
      @click="redirect('/water-rainfall')"
    />
    <Card
      img="https://thumbs.dreamstime.com/b/background-thunderstorm-illustration-clouds-rain-lightning-cartoon-155755077.jpg"
      icon="fa-solid fa-plug"
      detail="Relation of Electricity and Ranifall"
      name="Electic-Rain"
      @click="redirect('/elect-rainfall')"
    />
  </main>
</template>
