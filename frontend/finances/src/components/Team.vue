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
        <div><img src="@/assets/img/redact.svg" @click="editMember(index)"></div>
      </div>

    </div>

    <!-- кнопка -->
    <div class="create-button" @click="showModal=true">
      <button>Создать сотрудника</button>
    </div>
    <!-- модалка -->
    <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">
      <div class="modal">
        <h2>
          {{ editingIndex !== null ? 'Редактировать сотрудника' : 'Создать сотрудника' }}
        </h2>

        <label>
          Имя
          <input type="text" v-model="newMember.name" placeholder="Введите имя">
        </label>

        <label>
          Email
          <input type="email" v-model="newMember.email" placeholder="Введите email">
        </label>

        <label>
          Телефон
          <input type="text" v-model="newMember.phone" placeholder="+375...">
        </label>

        <label>
          Вид деятельности
          <input type="text" v-model="newMember.role" placeholder="Например: дизайнер">
        </label>

        <div class="modal-buttons">
          <button class="save-btn" @click="addMember">
            {{ editingIndex !== null ? 'Сохранить' : 'Создать' }}
          </button>
          <button class="cancel-btn" @click="showModal = false">Отмена</button>
        </div>
      </div>
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
      showModal: false,
      editingIndex: null, // какой элемент редактируем
      newMember: {
        name: '',
        email: '',
        phone: '',
        role: ''
      },
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
  },
  methods: {
    addMember() {
      if(this.newMember.name && this.newMember.email) {

        if(this.editingIndex !== null) {
          // ✏️ редактирование
          this.team[this.editingIndex] = { ...this.newMember };
        } else {
          // ➕ создание
          this.team.push({
            ...this.newMember,
            id: Date.now()
          });
        }

        // сброс
        this.newMember = {
          name: '',
          email: '',
          phone: '',
          role: ''
        };

        this.editingIndex = null;
        this.showModal = false;

      } else {
        alert('Введите имя и email');
      }
    },
    editMember(index) {
      this.editingIndex = index;

      // копируем данные в форму
      this.newMember = { ...this.team[index] };

      this.showModal = true;
    }
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.3);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
}

.modal {
  background: #fff;
  width: 500px;
  padding: 30px;
  border-radius: 15px;
}

.modal h2 {
  font-family: 'Manrope';
  font-size: 20px;
  margin-bottom: 20px;
}

.modal label {
  display: block;
  margin-bottom: 15px;
  font-size: 14px;
}

.modal input {
  width: 100%;
  height: 40px;
  margin-top: 5px;
  padding: 0 10px;
  border-radius: 6px;
  border: 1px solid #ccc;
  background: #F1F4FF;
}

.modal-buttons {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}

.save-btn {
  background: #3383FB;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 6px;
  cursor: pointer;
}

.cancel-btn {
  background: #E9E9E9;
  border: none;
  padding: 8px 16px;
  border-radius: 6px;
  cursor: pointer;
}

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