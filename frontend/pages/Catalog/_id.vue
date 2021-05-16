<script>
import gql from 'graphql-tag';

const Products_QUERY = gql`
    query($id: ID!) {
        Products(category_id: $id){
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
    apollo: {
        Products: {
            query: Products_QUERY,
            variables() {
                return{
                    id: this.$route.params.id
                }
            }
        }, 
    },
    methods:{
        ClickBy(Product){
          this.$router.push('/Buy/'+ Product.id)
        },
    },

}
</script>

<template>
    <div>
        <div class = " border-red-300 flex flex-col text-lg mb-60">
            <div class = "w-auto  flex items-center justify-between px-40 py-20 ">
                <div class = " h-80 w-72" v-for="(Product, idx) in Products" :key="idx"> 
                    <div class = "border-2 border-black w-full h-80"><img :src="`http://localhost:8080/static/`+Product.images[0].url" class="w-444 h-640" alt="если ты это видишь то что то с сервером"></div>
                    <button @click.prevent="ClickBy(Product)" class = "mt-6 w-full h-8 px-3 flex justify-around items-center bg-gray-600 text-white hover:text-gray-300">price: ${{Product.sub_products[0].price}}</button>
                </div>
            </div>
        </div>
    </div>
</template>