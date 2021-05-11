<template>
  <div>
<div class = "justify-center items-center w-full flex-col ">
    <div class = "w-auto flex justify-center items-center p-10 flex-row">
      <span class = "bg-black h-1 w-full mr-4"></span>
      <!-- <div class = "w-24 "><img src="" alt=""> сюда картинку </div> -->
      <span class = "bg-black h-1 w-full mr-4"></span>
    </div>
    <div class = " border-red-300 flex flex-col text-lg mb-60">
      <div class = "w-auto  flex items-center justify-between px-40 py-20 " v-for="category in Categories" :key="category">
        <div class = " h-80 w-72"> 
          <div class = "border-2 border-black w-full h-80">
            <img :src="category.image_url"
                alt="Placeholder"
                class="block h-auto w-full"
              />
          </div>
          <button @click.prevent="openCatalog(category)" class = "mt-6 w-full h-8 px-3 flex justify-around items-center bg-gray-600 text-white hover:text-gray-300">Open Category</button>
        </div>
      </div>
    </div>
    </div>
  </div>
</template>

<script>
import { gql, useLazyQuery  } from '@apollo/client';
const Categories_QUERY = gql`
    query {
      Categories{
        id
        name
        image_url
      }
    }
`
export default {
    layout: 'CatalogLayout',
    apollo: {
        Categories: {
            query: Categories_QUERY,
        }, 
    },   
    methods:{
        openCatalog(category){
          this.$router.push('/Catalog/'+ category.id)
        }
    } 
}

</script>