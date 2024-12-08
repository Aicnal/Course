<template>
  <div>
    <h2>课程列表</h2>
    <ul>
      <li v-for="course in courses" :key="course">{{ course }}</li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const courses = ref([]);

onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/courses');
    courses.value = response.data.courses;
  } catch (error) {
    console.error("获取课程失败:", error);
  }
});
</script>