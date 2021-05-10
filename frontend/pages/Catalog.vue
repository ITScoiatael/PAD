<script>

import { gql, useLazyQuery  } from '@apollo/client';

const Products_QUERY = gql`
    query {
        Products(category_id: "95809c10-d4e6-40d3-a695-1e6ef2589249"){
            id
            name
            description
            image_url
        }
    }
`


export default {
    apollo: {
        Products: {
            query: Products_QUERY,
        }, 
    },    
}

</script>

<template>
    <div>
        <div v-for="Product of Products" :key="Product" class="w-600 mx-auto rounded-lg bg-white border border-gray-200 p-5 text-gray-800 font-light mb-6">
            <div class="w-600 flex mb-4 items-center">
                <div class=" w-80 h-100 bg-gray-50 border border-gray-200">
                    <img :src="Product.image_url" class="w-444 h-640" alt="если ты это видишь то что то с сервером">
                </div>
                <div class="flex-grow pl-3">
                    <h4 class="font-bold text-sm uppercase text-gray-900">{{Product.name}}</h4>
                    <hr>
                    <h6 class="text-base font-bold text-sm uppercase text-gray-900">О товаре</h6>
                    <p>{{Product.description}}</p>
                </div>
            </div>
        </div>
    </div>
</template>