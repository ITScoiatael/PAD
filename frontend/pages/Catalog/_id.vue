<template>
  <!-- <div class="border-red-300 flex flex-col text-lg mb-60 items-center justify-around">
    <div class="w-auto lg:flex lg:items-center lg:justify-center py-20 flex-wrap">
      <div
        class="h-80 w-72 mx-20 flex items-center justify-center flex-col lg:mt-20 mt-24"
        v-for="(Product, idx) in Products"
        :key="idx"
      >
        <img
          :src="`http://localhost:8080/static/`+Product.images[0].url"
          class="border-2 border-black"
          alt="если ты это видишь то что то с сервером"
        />
        <button
          @click.prevent="ClickBy(Product)"
          class="mt-2 w-full h-8 px-3 flex justify-center items-center bg-gray-600 text-white hover:text-gray-300"
        >{{Product.name}}</button>
      </div>
    </div>
  </div> -->
<div class="div">
  <div class = "justify-center items-center w-full flex-col">
  <div class = "flex flex-col text-lg mb-60 items-center justify-around">
    <div class = " lg:flex lg:items-center lg:justify-center flex-wrap" >
        <!-- -->
        <div class = " mx-20 w-300 h-400 flex items-center justify-center flex-col lg:mt-20 mt-24 cursor-pointer" v-for="(Product, idx) in Products" :key="idx" @click.prevent="ClickBy(Product)"> 
          <div class = "w-full h-full overflow-hidden  overflow-hidden">
            <img class = "w-full h-full object-cover transform hover:scale-110 transition delay-150 duration-300 ease-in-out" :src="`http://localhost:8080/static/`+Product.images[0].url" alt="" >
          </div>
          <button class = "rounded-sm mt-2 flex flex-col justify-center items-center w-full bg-gray-600">
            <div class = "w-full h-full text-white">
            <div class = "text-2xl">price: ${{Product.sub_products[0].price}}</div>
            </div>
            <div class = "w-full justify-center items-center flex text-gray-400 "> 
              {{Product.name}}
            </div>
          </button>
        </div>
        <!-- -->
    </div>
  </div>
</div>

</div>

</template>

<script>
import gql from "graphql-tag";

const Products_QUERY = gql`
  query($id: ID!) {
    Products(category_id: $id) {
      id
      name
      description
      fabric
      manufacturer
      images {
        id
        url
      }
      sub_products {
        id
        price
        size
        color
        amount
      }
    }
  }
`;

export default {
  layout: "CatalogLayout",
  apollo: {
    Products: {
      query: Products_QUERY,
      variables() {
        return {
          id: this.$route.params.id
        };
      }
    }
  },
  methods: {
    ClickBy(Product) {
      this.$router.push("/buy/" + Product.id);
    }
  }
};
</script>