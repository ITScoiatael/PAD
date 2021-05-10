<template>
  <div>
   

    <div class="container my-12 mx-auto px-4 md:px-12 max-w-md mx-auto">
      <div class="flex flex-wrap -mx-7 lg:-mx-4">
        <div class="my-5 px-5 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/2">
          <article class="overflow-hidden rounded-lg shadow-lg">
            <NuxtLink to="Female_Catalog">
              <img
                alt="Placeholder"
                class="block h-auto w-full"
                src="http://localhost:8080/static/categories/female.png"
              />
            </NuxtLink>

            <header class="flex items-center justify-between leading-tight p-2 md:p-4">
              <h1 class="text-lg">
                <p class="text-black font-medium md:items-center" href="#">Female</p>
              </h1>
            </header>
          </article>
        </div>

        <div class="my-5 px-5 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/2">
          <article class="overflow-hidden rounded-lg shadow-lg">
            <NuxtLink to="Male_Catalog">
              <img
                alt="Placeholder"
                class="block h-auto w-full"
                src="http://localhost:8080/static/categories/male.png"
              />
            </NuxtLink>

            <header class="flex items-center justify-between leading-tight p-2 md:p-4">
              <h1 class="text-lg">
                <p class="text-black font-medium md:items-center" href="#">Male</p>
              </h1>
            </header>
          </article>
        </div>
      </div>
    </div>
  </div>
</template>