<template>
  <Header />
  <Nav page-title="Команда"/>

  <section class="team-page">

    <div class="team-container">

      <!-- header -->
      <div class="team-header grid">
        <div>Имя</div>
        <div>Email</div>
        <div>Телефон</div>
        <div>Вид деятельности</div>
        <div></div>
      </div>

      <!-- строки -->
      <div
        v-for="(member, index) in team"
        :key="index"
        class="team-row grid"
      >
        <div class="name">
          <input type="checkbox">
          {{ member.name }}
        </div>

        <div>{{ member.email }}</div>
        <div>{{ member.phone }}</div>
        <div>{{ member.role }}</div>
        <div><img src="@/assets/img/redact.svg"></div>
      </div>

    </div>

    <!-- кнопка -->
    <div class="create-button">
      <button>Создать сотрудника</button>
    </div>

  </section>

</template>

<script>
import Header from './Header.vue'
import Nav from './Nav.vue'

export default {
  components: {
    Header,
    Nav
  },

  data() {
    return {
      team: [
        {
          name: "Влада Насанович",
          email: "vlada@mail.com",
          phone: "+375291111111",
          role: "Веб-дизайнер"
        },
        {
          name: "Сергей Иванов",
          email: "serhei@mail.com",
          phone: "+375291111111",
          role: "Программист"
        },
        {
          name: "Анастасия Высоцкая",
          email: "anastasiya@mail.com",
          phone: "+375291111111",
          role: "Копирайтер"
        },
        {
          name: "Анатолий Петров",
          email: "anatoliy@mail.com",
          phone: "+375291111111",
          role: "Рилсмейкер"
        }
      ]
    }
  }
}
</script>

<style scoped>

.team-page{
  padding: 40px;
  background:#F1F4FF;
  min-height:100vh;
}

/* контейнер таблицы */

.team-container{
  width:100%;
  max-width:940px;
  margin:auto;

  background:white;
  border:1px solid #BBBBBB;
  border-radius:15px;
  overflow:hidden;
}

/* grid */

.grid{
  display:grid;
  grid-template-columns:
  2fr
  2fr
  1.5fr
  1.5fr
  40px;

  align-items:center;
}

/* header */

.team-header{
  background:#DBE1F6;
  padding:16px 20px;

  font-family:Manrope;
  font-size:14px;
  font-weight:500;

  border-bottom:1px solid #BBBBBB;
}

/* строки */

.team-row{
  padding:18px 20px;

  font-family:Manrope;
  font-size:14px;
  font-weight:300;

  border-bottom:1px solid #BBBBBB;
}

.team-row:last-child{
  border-bottom:none;
}

/* имя + чекбокс */

.name{
  display:flex;
  align-items:center;
  gap:10px;
}

/* кнопка */

.create-button{
  width:100%;
  max-width:940px;
  margin:20px auto 0;
}

.create-button button{

  width:100%;
  height:40px;

  border:none;
  border-radius:15px;

  background:#3383FB;
  color:white;

  font-family:Manrope;
  font-size:14px;

  cursor:pointer;
}

.create-button button:hover{
  background:#2c72da;
}

</style>