<script>
import { gql, useLazyQuery  } from '@apollo/client';

function GetDataFromLS(){
    if(process.browser){
        return localStorage.getItem("CategoryqwertyId");
    }
    return null
}

const categoryid_FromLS = GetDataFromLS()
console.log(categoryid_FromLS)

const Products_QUERY = gql`
    query($id: ID!) {
        Products(category_id: $id){
            id
            name
            description
            image_url
        }
    }
`
export default {
    layout: 'CatalogLayout',
    data(){
        return{
            category_id: "5bdec776-9518-47b7-8a45-925ae1aa2c9c"

            // I don't know why this shit doesn't want to work))
            //category_id: categoryid_FromLS
        }  
    },
    apollo: {
        Products: {
            query: Products_QUERY,
            variables() {
                return{
                    id: this.category_id
                }
            }
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