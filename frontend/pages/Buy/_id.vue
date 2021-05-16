
<script>
import gql from 'graphql-tag';

const SubProduct_QUERY = gql`
    query ($id: ID!) {
        Product(id: $id){
            id
            name
            description
            fabric
            manufacturer
            images{
              id
              url
            }
            sub_products{
                id
                price
                size
                color
                amount
            }
        }
    }
`



export default {
    layout: 'CatalogLayout',
    data(){
      return{
        index: 0
      }
    },
    methods:{
      incremectIndex(){
        if(this.Product.sub_products.length - 1 > this.index)
          this.index++
      },
      decrenemtIndex(){
        if(0 < this.index)
          this.index--
      }
    },
    apollo: {
        Product: {
            query: SubProduct_QUERY,
            variables() {
                return{
                    id: this.$route.params.id
                }
            }
        },
    },

}
</script>

<template>
  <div class="flex w-full flex-col h-full">

    <LineWithLogo/>
    <button @click="incremectIndex()">go forvart</button>
    <button @click="decrenemtIndex()">go back</button>
    <div class="w-auto flex justify-center items-center flex-col p-6 py-6">
      <div class="w-full flex justify-center items-center mb-5">
        <div class="w-20 h-20 border-2 border-black mr-2">
          <img :src="`http://localhost:8080/static/`+Product.images[0].url" class="w-444 h-640" alt="если ты это видишь то что то с сервером">
          </div>
        <div
          class="w-full border-2 border-black h-20 bg-black rounded-md flex justify-center items-center flex-row text-white text-xl"
        >
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Name:</div>
            <div class="ml-2">{{Product.name}}</div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Color:</div>
            <div class="border-цршеу-300 border-2 h-4 w-4 ml-2 bg-yellow-300 rounded-sm">{{Product.sub_products[index].color}}</div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Size:</div>
            <div class="ml-2">{{Product.sub_products[index].size}}</div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Material:</div>
            <div class="ml-2">{{Product.fabric}}</div>
          </div>
          <div class="mr-10 flex flex-row items-center">
            <div class="text-gray-400">Made:</div>
            <div class="ml-2">{{Product.manufacturer}}</div>
          </div>
          <div class="flex flex-row items-center">
            <div class="text-gray-400">Price:</div>
            <div class="ml-2">${{Product.sub_products[index].price}}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>