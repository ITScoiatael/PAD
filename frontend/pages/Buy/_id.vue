<template>  
 <div class = "w-full flex justify-center items-center flex-row px-10 py-20 ">
   <div class = "w-full flex justify-center items-center flex-row ">
     <div class = "border-2 border-black w-300 h-330">
       <img src="" alt="photo" class = "w-full h-full">
     </div>
     <div class = "flex flex-col mx-3 justify-between items-center mr-40">
       <div class = "mb-2 w-20 h-20 shadow-md bg-gray-300" v-for="(productImage, index) in Product.images" :key="index">
         <img :src="`http://localhost:8080/static/`+productImage.url" alt= "цвет товара">
       </div>
     </div>
     <div class = "bg-gray-300 shadow-md flex flex-col justify-center items-center rounded-md w-full px-5 ">
       <div class = "w-full mb-3 mt-3 flex flex-row">
         <div class = "rounded-md text-white bg-black w-60 flex flex-col justify-center items-center  text-2xl">
           <div class = "flex justify-start items-start ml-2 w-full h-9">Price</div>
           <div class = "flex justify-start items-start ml-2 w-full h-9">Color</div>
           <div class = "flex justify-start items-start ml-2 w-full h-9">Size</div>
           <div class = "flex justify-start items-start ml-2 w-full h-9">Material</div>
           <div class = "flex justify-start items-start ml-2 w-full h-9">Made In</div>
           <div class = "flex justify-start items-start ml-2 w-full h-9">Amount</div>
         </div>
         <div class = "w-full ">
           <div class = "w-full ml-3 h-9 flex items-center">
             <div class = "text-2xl font-medium pb-2">${{Product.sub_products[0].price}}</div>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center pb-1">
              <button class = "w-8 h-8 border-black border-2 rounded-md bg-white focus:bg-gray-400 focus:outline-none" v-for="(sub_product, idx) in Product.sub_products" :key="idx" >
                {{sub_product.color}}
                </button>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center pt-1">
             <button class = "w-8 h-8 border-black border-2 rounded-md focus:bg-black focus:text-white hover:bg-black hover:text-white" v-for="(size, index) in Sizes" :key="index">{{size}}</button>
            </div>
           <div class = "w-full ml-3 h-9 flex items-center">
             <div class = "text-2xl font-medium">{{Product.fabric}}</div>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center">
             <div class = "text-2xl font-medium">{{Product.manufacturer}}</div>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center pt-1">
             <div class = "w-auto mr-2 flex items-center justify-center">
              <button @click.prevent="decrenemtAmount()" class="w-5 h-5 outline-none rounded-3xl flex items-center justify-center text-white bg-gray-700 text-md hover:bg-gray-800 focus:outline-none">-</button>
             </div>
             <div class = "w-20 bg-white h-9 rounded-md flex justify-center items-center text-xl overflow-x-auto">{{Amount}}</div>
             <div class = "w-auto ml-2">
              <button @click.prevent="incremectAmount()" class="w-5 h-5 outline-none rounded-3xl flex items-center justify-center text-white bg-gray-700 text-md hover:bg-gray-800 focus:outline-none">+</button>
             </div>
           </div>
         </div>
         <div class = "w-full py-4 flex justify-center items-center text-white ">
           <div class = "w-300 h-full bg-black flex flex-col rounded-xl justify-start items-center px-4 py-2">
             <div class = "w-full text-2xl">About product</div>
             <div class = " w-full h-full text-xl flex items-center pb-8">{{Product.description}}</div>
           </div>
         </div>
       </div>
       <div class = "w-full mb-3 px-20 flex flex-row rounded-md bg-black py-2">
         <div class = " text-white flex items-center justify-center flex-col">
           <div class = "text-2xl tracking-widest">USA</div>
           <div class = "text-2xl tracking-widest">Russia</div>
           <div class = "text-2xl tracking-widest">Spain</div>
         </div>
        <span class = "ml-20 mr-20 h-100 bg-white w-1 mb-2 mt-2"></span>
        <div class = "w-full flex flex-col">
          <div class = "border-2 border-red-300 w-full h-full"></div>
          <div class = "border-2 border-red-300 w-full h-full"></div>
          <div class = "border-2 border-red-300 w-full h-full"></div>
        </div>
       </div>
       <div class = "w-full mb-3 flex flex-row text-white">
         <div class = "w-full mr-80">
           <button @click.prevent="goBack()" class = "uppercase flex justify-center items-center w-full bg-black rounded-xl hover:bg-gray-900 py-2 focus:outline-none">Cancel</button>
         </div>
         <div class = "w-full">
           <button class = "uppercase flex justify-center items-center w-full bg-black rounded-xl hover:bg-gray-900 py-2 focus:outline-none ">buy</button>
         </div>
       </div>
     </div>
   </div>
 </div>
</template>

<script>
import gql from 'graphql-tag';

const Product_QUERY = gql`
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

function unique(arr) {
  let result = [];

  for (let str of arr) {
    if (!result.includes(str)) {
      result.push(str);
    }
  }

  return result;
}

export default {
    layout: 'CatalogLayout',
    data(){
      return{
        Amount: 0,
      }
    },
    computed: {
      Sizes: function (){
        let temp = []
        for(let i = 0; i < this.Product.sub_products.length - 1; i++){
          temp[i] = this.Product.sub_products[i].size
        }
        return unique(temp)
      }
    },
    mounted() { },
    methods:{
      goBack(){
        this.$router.push("/");
      },
      incremectAmount(){
        if(this.Product.sub_products.length - 1 > this.Amount)
          this.Amount++
      },
      decrenemtAmount(){
        if(0 < this.Amount)
          this.Amount--
      },
      
    },
    apollo: {
        Product: {
            query: Product_QUERY,
            variables() {
                return{
                    id: this.$route.params.id
                }
            }
        },
    },
}
</script>