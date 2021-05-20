<template>
  <div class="min-h-screen bg-GrayImage">
    <div class="w-full h-screen flex justify-center items-center">
      <div
        class="rounded-3xl bg-black w-700 h-400 flex justify-center items-center flex-col"
      >
        <div
          class="mt-2 w-full justify-center items-center flex-col flex px-10"
        >
          <p class="text-gray-50 text-2xl mb-2">Login</p>
          <input
            type="text"
            class="outline-none pl-1 h-14 w-full rounded-md focus:placeholder-transparent"
            placeholder="Your Login"
          />
        </div>
        <div
          class="mt-14 w-full justify-center items-center flex-col flex px-10"
        >
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
          >
            Enter
          </button>
        </div>
      </div>
    </div>
  </div>
</template>