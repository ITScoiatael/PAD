<script>
import { gql, useLazyQuery } from "@apollo/client";
const Categories_QUERY = gql`
  query {
    Categories {
      id
      name
      image_url
    }
  }
`;
export default {
  layout: "CatalogLayout",
  apollo: {
    Categories: {
      query: Categories_QUERY
    }
  },
  methods: {
    openCatalog(category) {
      localStorage.setItem("CategoryqwertyId", category.id);
      this.$router.push("/Catalog/" + category.name);
    }
  }
};
</script>


<template>
  <div>
    <div class="container my-12 mx-auto px-4 md:px-12 max-w-md mx-auto">
      <div class="flex flex-wrap -mx-7 lg:-mx-4">
        <div
          class="my-5 px-5 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/2 cursor-pointer"
          v-for="category in Categories" :key="category" >
          <article
            class="overflow-hidden rounded-lg shadow-lg"
            @click.prevent="openCatalog(category)"
          >
            <img :src="category.image_url" alt="Placeholder" class="block h-auto w-full" />

            <header class="flex items-center justify-between leading-tight p-4 md:p-4">
              <h1 class="text-lg">
                <p class="text-black font-medium md:items-center">{{category.name}}</p>
              </h1>
            </header>
          </article>
        </div>
      </div>
    </div>

    <div class="justify-center items-center w-full flex-col">
      <div class="w-auto flex justify-center items-center p-10 flex-row">
        <span class="bg-black h-1 w-full mr-4"></span>
        <div class="w-24">
          <img src alt />
          <img src="../../images/square.png" alt />
        </div>
        <span class="bg-black h-1 w-full mr-4"></span>
      </div>
    </div>
  </div>
</template>