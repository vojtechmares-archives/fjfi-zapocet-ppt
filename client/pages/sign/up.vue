<template>
  <div class="sign-in">

  <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
    <h3 class="font-bold text-gray-900 text-xl text-center">Sign Up</h3>
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
        Username
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="username" type="text" placeholder="Username">
    </div>
    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
        Password
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="******************">
    </div>
    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
        Password Again
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="******************">
    </div>
    <div class="flex items-center justify-between">
      <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="button" @click="submit">
        Sign Up
      </button>
    </div>
  </form>
  <p class="text-center text-gray-500 text-xs">
    Zápočtový program pro předmět PPT na FJFI ČVUT &copy;&nbsp;2020 Vojtěch Mareš
  </p>
</div>
</template>

<script>
export default {
  methods: {
    submit() {
      this.signUp(username, password);
    },
    signUp(username, password) {
      // POST /api/v1/users/sign-up
    }
  }
}
</script>

<style>
body {
  background-color: rgb(237, 242, 247);
}

.sign-in {
  margin: 0 auto;
  margin-top: 10rem;
  max-width: 20rem;
}
</style>
