<template>
  <div class="chart-container">
    <svg :width="width" :height="height" viewBox="0 0 100 50">
      <!-- Фон -->
      <rect x="0" y="0" width="100" height="50" fill="#F1F4FF" rx="4" />
      
      <!-- Синяя линия выручки -->
      <path
        :d="linePath"
        fill="none"
        stroke="#3383FB"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        :stroke-dasharray="pathLength"
        :stroke-dashoffset="pathLength"
        ref="line"
      />
    </svg>
  </div>
</template>

<script>
import { ref, onMounted, watch, computed } from "vue";

export default {
  name: "Revenue",
  props: {
    data: {
      type: Array,
      required: true, // ожидаем массив чисел [50, 80, 30...]
    },
    width: {
      type: Number,
      default: 500,
    },
    height: {
      type: Number,
      default: 200,
    },
  },
  setup(props) {
    const line = ref(null);

    const linePath = computed(() => {
      if (!props.data.length) return "";
      const n = props.data.length;
      const maxVal = Math.max(...props.data);
      const minVal = Math.min(...props.data);
      const scaleY = (v) => ((maxVal - v) / (maxVal - minVal || 1)) * 50;
      const scaleX = (i) => (i / (n - 1)) * 100;

      return props.data
        .map((v, i) => `${i === 0 ? "M" : "L"} ${scaleX(i)} ${scaleY(v)}`)
        .join(" ");
    });

    const pathLength = ref(0);

    onMounted(() => {
      pathLength.value = line.value.getTotalLength();
      // Анимация линии
      line.value.style.transition = "stroke-dashoffset 1.2s ease-out";
      setTimeout(() => {
        line.value.style.strokeDashoffset = "0";
      }, 50);
    });

    watch(
      () => props.data,
      () => {
        // при обновлении данных пересчитываем путь
        pathLength.value = line.value.getTotalLength();
        line.value.style.strokeDasharray = pathLength.value;
        line.value.style.strokeDashoffset = pathLength.value;
        setTimeout(() => {
          line.value.style.strokeDashoffset = "0";
        }, 50);
      }
    );

    return { line, linePath, pathLength };
  },
};
</script>

<style scoped>
.chart-container {
  width: 100%;
  height: 100%;
}
svg {
  width: 100%;
  height: 100%;
}
</style>