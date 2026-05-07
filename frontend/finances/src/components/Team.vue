<template>
  <Header />
  <Nav page-title="Команда" class="nav-component"/>

  <section class="team-page">
    
    <!-- Заголовок страницы -->
    <h1 class="page-title">Команда</h1>

    <div class="team-container">

      <!-- Кнопка для мобильных -->
      <button class="mobile-add-btn" @click="showModal=true">
        <span>+</span>
      </button>

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
          <span class="member-name">{{ member.name }}</span>
        </div>

        <div class="member-email">{{ member.email }}</div>
        <div class="member-phone">{{ member.phone }}</div>
        <div class="member-role">{{ member.role }}</div>
        <div><img src="@/assets/img/redact.svg" @click="editMember(index)" class="edit-icon"></div>
      </div>

    </div>

    <!-- кнопка для десктопа -->
    <div class="create-button desktop-btn" @click="showModal=true">
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
      editingIndex: null,
      newMember: {
        name: '',
        email: '',
        phone: '',
        role: ''
      },
      team: []
    }
  },
  methods: {
    // Замените метод fetchTeamMembers
    async fetchTeamMembers() {
      try {
        const token = localStorage.getItem('token')
        if (!token) {
          console.error('No token found')
          return
        }
        
        // Убираем credentials: 'include' и Content-Type для GET запроса
        const response = await fetch('http://localhost:8080/api/team', {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${token}`,
            'Accept': 'application/json'
          }
        })
        
        if (!response.ok) {
          if (response.status === 401) {
            localStorage.removeItem('token')
            window.location.href = '/login'
            return
          }
          throw new Error(`HTTP error! status: ${response.status}`)
        }
        
        const data = await response.json()
        this.team = data.members || []
      } catch (error) {
        console.error('Error fetching team members:', error)
      }
    },
    
      async addMember() {
      if(this.newMember.name && this.newMember.email) {
        try {
          const token = localStorage.getItem('token')
          if (!token) {
            window.location.href = '/login'
            return
          }
          
          let response;
          
          if(this.editingIndex !== null) {
            const member = this.team[this.editingIndex];
            response = await fetch(`http://localhost:8080/api/team/${member.id}`, {
              method: 'PUT',
              headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
                'Accept': 'application/json'
              },
              body: JSON.stringify(this.newMember)
            });
          } else {
            response = await fetch('http://localhost:8080/api/team', {
              method: 'POST',
              headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
                'Accept': 'application/json'
              },
              body: JSON.stringify(this.newMember)
            });
          }
          
          if (response.ok) {
            await this.fetchTeamMembers();
            
            this.newMember = {
              name: '',
              email: '',
              phone: '',
              role: ''
            };
            this.editingIndex = null;
            this.showModal = false;
          } else {
            const error = await response.json();
            alert(error.error || 'Ошибка при сохранении');
          }
        } catch (error) {
          console.error('Error saving team member:', error);
          alert('Ошибка при сохранении');
        }
      } else {
        alert('Введите имя и email');
      }
    },
    
    editMember(index) {
      this.editingIndex = index;
      this.newMember = { ...this.team[index] };
      this.showModal = true;
    },
    
    async deleteMember(index) {
      if (confirm('Вы уверены, что хотите удалить этого сотрудника?')) {
        try {
          const token = localStorage.getItem('token')
          const member = this.team[index];
          
          const response = await fetch(`/api/team/${member.id}`, {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${token}`
            }
          });
          
          if (response.ok) {
            this.team.splice(index, 1);
          } else {
            const error = await response.json();
            alert(error.error || 'Ошибка при удалении');
          }
        } catch (error) {
          console.error('Error deleting team member:', error);
          alert('Ошибка при удалении');
        }
      }
    }
  },
  
  mounted() {
    this.fetchTeamMembers();
  }
}
</script>

<style scoped>
* {
  box-sizing: border-box;
}

/* Заголовок страницы */
.page-title {
  text-align: center;
  font-family: "Manrope";
  font-weight: 500;
  font-size: 30px;
  line-height: 140%;
  letter-spacing: 0%;
  color: #000000;
  margin-bottom: 30px;
}

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
  width: 90%;
  max-width: 500px;
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
  position: relative;
}

/* Кнопка для мобильных */
.mobile-add-btn {
  display: none;
  position: absolute;
  top: 10px;
  right: 10px;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #3383FB;
  color: white;
  border: none;
  font-size: 24px;
  cursor: pointer;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2);
  align-items: center;
  justify-content: center;
}

.mobile-add-btn span {
  line-height: 1;
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

.edit-icon {
  cursor: pointer;
}

/* кнопка для десктопа */

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

/* Адаптивность для мобильных устройств */
@media (max-width: 768px) {
  /* Скрываем Nav */
  .nav-component {
    display: none !important;
  }
  
  /* Адаптируем заголовок */
  .page-title {
    font-size: 24px;
    margin-bottom: 20px;
  }
  
  /* Адаптируем страницу */
  .team-page {
    padding: 20px 10px;
  }
  
  /* Показываем мобильную кнопку */
  .mobile-add-btn {
    display: flex;
  }
  
  /* Скрываем десктопную кнопку */
  .desktop-btn {
    display: none;
  }
  
  /* Адаптируем контейнер */
  .team-container {
    width: 100%;
    max-width: 100%;
    margin: 0;
    border-radius: 10px;
  }
  
  /* Адаптируем grid */
  .grid {
    grid-template-columns: 1.5fr 2fr 1.5fr 1fr 30px;
    gap: 5px;
    font-size: 12px;
  }
  
  /* Адаптируем header */
  .team-header {
    padding: 12px 10px;
    font-size: 11px;
  }
  
  /* Адаптируем строки */
  .team-row {
    padding: 12px 10px;
    font-size: 11px;
  }
  
  /* Уменьшаем чекбокс */
  .name input[type="checkbox"] {
    width: 16px;
    height: 16px;
  }
  
  .name {
    gap: 5px;
  }
  
  /* Адаптируем иконку редактирования */
  .edit-icon {
    width: 16px;
    height: 16px;
  }
  
  /* Адаптируем модальное окно */
  .modal {
    width: 95%;
    padding: 20px;
    margin: 10px;
  }
  
  .modal h2 {
    font-size: 18px;
  }
  
  .modal input {
    height: 36px;
    font-size: 14px;
  }
}

@media (max-width: 480px) {
  /* Адаптируем заголовок для маленьких экранов */
  .page-title {
    font-size: 20px;
    margin-bottom: 15px;
  }
  
  .team-page {
    padding: 15px 5px;
  }
  
  /* Еще сильнее сжимаем grid для маленьких экранов */
  .grid {
    grid-template-columns: 1.2fr 1.8fr 1.2fr 0.8fr 25px;
    gap: 3px;
    font-size: 10px;
  }
  
  .team-header {
    padding: 10px 8px;
    font-size: 10px;
  }
  
  .team-row {
    padding: 10px 8px;
    font-size: 10px;
  }
  
  .name {
    gap: 3px;
  }
  
  .name input[type="checkbox"] {
    width: 14px;
    height: 14px;
  }
  
  .mobile-add-btn {
    width: 36px;
    height: 36px;
    font-size: 20px;
    top: 8px;
    right: 8px;
  }
  
  .modal {
    padding: 15px;
  }
  
  .modal-buttons {
    flex-direction: column;
  }
  
  .save-btn, 
  .cancel-btn {
    width: 100%;
  }
}
</style>