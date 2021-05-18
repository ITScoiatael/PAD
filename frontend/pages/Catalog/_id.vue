<template>
  <div class="border-red-300 flex flex-col text-lg mb-60 items-center justify-around">
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
        >price: ${{Product.sub_products[0].price}}</button>
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
      this.$router.push("/Buy/" + Product.id);
    }
  }
};
</script>