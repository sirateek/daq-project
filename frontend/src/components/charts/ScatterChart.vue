<template>
  <div class="box">
    <p class="grid text-black text-sm z-50 w-full place-content-center rotate-[270deg]">{{ yAxisLabel }}</p>
    <Scatter :chart-options="{
      responsive: true,
    }" :chart-data="{
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
import { Scatter } from "vue-chartjs";
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
  name: "ScatterChart Component",
  components: {
    Scatter,
  },
  props: {
    labels: {
      type: Object as PropType<String>,
      required: true,
    },
    data: {
      type: Object as PropType<Array<number>>,
      required: true,
    },
    dataLabel: {
      type: String,
      default: "Scatter",
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
  