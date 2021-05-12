<template>
  <div class="flex w-full flex-col h-full">
    
    <Header/>

    <div class="w-auto flex justify-center items-center p-10 flex-row">
      <span class="bg-black h-1 w-full mr-4"></span>
      <div class="w-24">
        <img src alt /> сюда картинку
      </div>
      <span class="bg-black h-1 w-full mr-4"></span>
    </div>

    <div class="w-auto flex justify-center items-center flex-col p-6 py-6">
      <div class="w-full flex justify-center items-center mb-5">
        <div class="w-20 h-20 border-2 border-black mr-2">фото товара</div>
        <div
          class="w-full border-2 border-black h-20 bg-black rounded-md flex justify-center items-center flex-row text-white text-xl"
        >
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Name:</div>
            <div class="ml-2">Тряпка</div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Color:</div>
            <div class="border-цршеу-300 border-2 h-4 w-4 ml-2 bg-yellow-300 rounded-sm"></div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Size:</div>
            <div class="ml-2">45</div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Material:</div>
            <div class="ml-2">Cotton</div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Made:</div>
            <div class="ml-2">China</div>
          </div>
          <div class="flex flex-row items-center">
            <div class="text-gray-400">Price:</div>
            <div class="ml-2">3.8$</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>