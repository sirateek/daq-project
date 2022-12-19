<template>
  <ul class="cards">
    <li>
      <a class="card">
        <img :src="img" class="card__image" alt="" />
        <div class="card__overlay">
          <div class="card__header bg-gray-400 bg-opacity-60">
            <svg class="card__arc" xmlns="http://www.w3.org/2000/svg">
            </svg>
            <i :class="icon"></i>
            <div class="card__header-text">
              <h3 class="card__title">{{name}}</h3>
            </div>
          </div>
          <p class="card__description text-white bg-gray-400 bg-opacity-60">
            {{detail}}
          </p>
        </div>
      </a>
    </li>
  </ul>
</template>

<script>
export default {
  name: "card",
  props: {
    img: String,
    icon: String,
    detail: String,
    name: String
  }
} 
</script>

<style scoped>
.cards {
  display: grid;
  height: 100%;
}

.card {
  position: relative;
  display: block;
  height: 100%;
  overflow: hidden;
  text-decoration: none;
  border-radius: 12px;
}

.card__image {
  width: 100%;
  height: 100%;
}

.card__overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  transform: translateY(100%);
  transition: 0.2s ease-in-out;
}

.card:hover .card__overlay {
  transform: translateY(0);
}

.card__header {
  position: relative;
  display: flex;
  align-items: center;
  gap: 2em;
  padding: 2em;
}

.card__arc {
  width: 80px;
  height: 80px;
  position: absolute;
  bottom: 100%;
  right: 0;
}

.card:hover .card__header {
  transform: translateY(0);
}



.card__title {
  font-size: 1em;
  margin: 0 0 0.3em;
  color: white;
}


.card__description {
  padding: 0 2em 2em;
  margin: 0;
}
</style>
