<template>
  <!-- <section class="w-full h-screen flex  flex-col justify-center items-center">
        <div class="form">
            <h1>Login</h1>
            <form action="" class="flex flex-col items-start my-2">
                <input type="text" class="w-96 h-10 my-2">
                <input type="text" class="w-96 h-10 my-2">
                <button type="submit" class="my-2 bg-blue-500 py-2 px-3">Login!</button>
            </form>
        </div>
  </section>-->

  <div class="w-full h-screen flex justify-center items-center">
    <div class="rounded-3xl bg-black w-700 h-400 flex justify-center items-center flex-col">
      <div class="mt-2 w-full justify-center items-center flex-col flex px-10">
        <p class="text-gray-50 text-2xl mb-2">Login</p>
        <input
          type="text"
          class="outline-none pl-1 h-14 w-full rounded-md focus:placeholder-transparent"
          placeholder="Your Login"
        />
      </div>
      <div class="mt-14 w-full justify-center items-center flex-col flex px-10">
        <p class="text-gray-50 text-2xl mb-2">Password</p>
        <input
          type="password"
          class="w-full outline-none pl-1 h-14 rounded-md focus:placeholder-transparent"
          placeholder="Your Password"
        />
      </div>
      <div class="flex justify-center items-center mt-10">
        <button
          class="bg-white my-2 px-4 py-1 text-2xl outline-none focus:bg-gray-300 focus:outline-none rounded-md"
        >Enter</button>
      </div>
    </div>
  </div>
</template>


<script>
export default {};
</script>