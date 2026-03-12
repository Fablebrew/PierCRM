<template>
  <section id="nav-section">
    <div class="breadcrumb">
      <h1>{{ pageTitle }}</h1>
    </div>
    <div class="nav">
      <ul>
        <li><a href="#">Проекты</a></li>
        <li><a href="#">Прочие расходы</a></li>
        <li><a href="#">Команда</a></li>
      </ul> 
    </div>
  </section>
</template>

<script>
export default {
  name: "Nav",
  props: {
    pageTitle: {
      type: String,
      default: "Рабочий стол" // если не передан prop
    }
  }
};
</script>

<style scoped>
#nav-section {
  border-bottom: 1px solid #BBBBBB;
  display: flex;
  padding: 0.5% 5% 0.5% 5%;
  align-items: center;
  position: relative;
  background-color: #F1F4FF;
}
.breadcrumb h1 {
  font-family: "Manrope";
  font-weight: 400;
  font-style: Regular;
  font-size: 22px;
  line-height: 140%;
  letter-spacing: 0%;
  text-align: center;
  color: #000000B2;
}
.nav {
  align-content: center;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}
.nav ul {
  list-style: none;
  display: flex;
}
.nav ul li {
  margin-left: 20px;
}
.nav ul li a {
  text-decoration: none;
  font-family: "Manrope";
  font-weight: 400;
  font-style: Regular;
  font-size: 18px;
  line-height: 140%;
  letter-spacing: 0%;
  text-align: center;
  color: #000000B2;
}
</style>