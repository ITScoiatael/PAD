<template>
  <div>
    <div class="container my-12 mx-auto px-4 md:px-12 max-w-md mx-auto">
      <div class="flex flex-wrap -mx-7 lg:-mx-4">
        <div v-for="category in Categories" :key="category" class="my-5 px-5 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/2">
          <a @click.prevent="openCatalog(category)">
            <article class="overflow-hidden rounded-lg shadow-lg">
              <img :src="category.image_url"
                alt="Placeholder"
                class="block h-auto w-full"
              />

              <header class="flex items-center justify-between leading-tight p-2 md:p-4">
                <h1 class="text-lg">
                  <p class="text-black font-medium md:items-center" href="#">{{category.name}}</p>
                </h1>
              </header>
            </article>
          </a>
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