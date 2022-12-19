<template>
  <div class="box">
    <p class="grid text-black text-sm z-50 w-full place-content-center rotate-[270deg]">{{ yAxisLabel }}</p>
    <Line :chart-options="{
      responsive: true,
    }" :chart-data="{
  labels: labels,
  datasets: [
    {
      data: data,
      label: dataLabel,
      backgroundColor: dataColor,
      borderColor: dataColor,
    },
  ],
}" :width="400" :height="150" />
  </div>
  <p class="text-black text-sm z-50 w-full text-center">{{ xAxisLabel }}</p>
</template>

<script lang="ts">
import { PropType } from "vue";
import { Line } from "vue-chartjs";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  LineElement,
  CategoryScale,
  LinearScale,
  PointElement
} from "chart.js";

ChartJS.register(
  Title,
  Tooltip,
  Legend,
  LineElement,
  CategoryScale,
  LinearScale,
  PointElement
);

export default {
  name: "LineChart Component",
  components: {
    Line,
  },
  props: {
    labels: {
      type: Object as PropType<Array<number>>,
      required: true,
    },
    data: {
      type: Object as PropType<Array<number>>,
      required: true,
    },
    dataLabel: {
      type: String,
      default: "Line",
    },
    dataColor: {
      type: String,
      default: "#f87979",
    },
    xAxisLabel: {
      type: String,
      required: true
    },
    yAxisLabel: {
      type: String,
      required: true
    }
  },
};
</script>

<style>
.box {
  display: grid;
  grid-template-columns: 3% 97%;
}
</style>