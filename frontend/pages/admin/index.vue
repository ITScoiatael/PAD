<template>
  <section class="flex w-full flex-col h-full">

    <div class="w-full flex justify-center items-center">
      <div class="h-200 w-200 border-2 flex justify-center items-center flex-col mr-2">
        <img src="" alt="ссылка на фото" class="w-full h-full object-cover">
      </div>
      <div class="h-200 w-200 border-2 flex justify-center items-center flex-col  mr-2">
        <img src="" alt="ссылка на фото" class="w-full h-full object-cover">
      </div>
      <div class="h-200 w-200 border-2 flex justify-center items-center flex-col  mr-2">
        <img src="" alt="ссылка на фото" class="w-full h-full object-cover">
      </div>
      <div class="h-200 w-200 border-2 flex justify-center items-center flex-col mr-2">
        <img src="" alt="ссылка на фото" class="w-full h-full object-cover">
      </div>
      <div class=" h-200 w-200 border-2 flex justify-center items-center flex-col">
        <img src="" alt="ссылка на фото" class="w-full h-full object-cover">
      </div>
    </div>
    <div class="w-full mt-20 ">


      <div class="w-full flex flex-row mb-4">
        <div class="flex flex-col  px-10">
          <div class="text-black uppercase font-bold">Count</div>
          <button
            class="w-10 h-10 border-2 border-black focus:outline-none hover:bg-gray-200 mt-2 rounded-md items-center justify-center">+</button>
        </div>
        <textarea class="rounded-md resize border-2 border-black p-1"></textarea>
      </div>


      <div class="w-full flex flex-row mb-4">
        <div class="flex flex-col  px-10 mr-4">
          <div class="text-black uppercase font-bold h-full justify-start flex items-start">size</div>
        </div>
        <div class="flex flex-col justify-center items-center">
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
        </div>
      </div>

      <div class="w-full flex flex-row mb-4">
        <div class="flex flex-col  px-10 mr-4">
          <div class="text-black uppercase font-bold h-full justify-start flex items-start">color</div>
        </div>
        <div class="flex flex-col justify-center items-center">
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
          <textarea class="rounded-md w-20 resize border-2 border-black h-16 mb-1 p-1"></textarea>
        </div>
      </div>


      <div class="w-full flex flex-row mb-4">
        <div class="flex flex-col  px-10">
          <div class="text-black uppercase font-bold h-full justify-center flex items-center">Производитель</div>
        </div>
        <textarea class="rounded-md resize border-2 border-black p-2"></textarea>
      </div>

      <div class="w-full flex flex-row mb-4">
        <div class="flex flex-col  px-10">
          <div class="text-black uppercase font-bold h-full justify-center flex items-center">цена</div>
        </div>
        <textarea class="rounded-md resize border-2 border-black p-2"></textarea>
      </div>

      <div class="w-full mb-10 mt-10 flex flex-col items-center justify-center px-20 mb-10">
        <div class="w-full flex justify-center">
          <textarea class="rounded-md w-full resize border-2 border-black h-20 p-2"
            placeholder="Information"></textarea>
        </div>
        <div class="mt-5 flex flex-row">
          <button
            class="px-6 py-1 bg-gray-500 text-white hover:bg-gray-400 w-full rounded-md uppercase font-bold">Enter</button>
        </div>
      </div>
    </div>
  </section>
</template>
