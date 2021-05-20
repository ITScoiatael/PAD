<template>  
 <div class = "px-8 pb-20" v-if="Product">
   <div class = "w-full flex justify-center items-center 2xl:flex-row flex-col">
     <div class = "border-2 border-black w-300 h-330">
       <img :src="`http://localhost:8080/static/`+Product.images[0].url" id="MainImage" alt="photo" class = "w-full h-full">
     </div>
     <div class = "flex 2xl:flex-col flex-row 2xl:mx-3 mx-3 justify-between items-center">
       <button class = " 2xl:mx-3 mr-3 my-3 w-20 h-20 shadow-md bg-gray-300" v-for="(image, idx) in Images" :key="idx" @click.prevent="SelectImage(idx)"> 
         <img :src="`http://localhost:8080/static/`+image.url" alt= "ФОТО товара">
       </button> 
     </div>
     <div class = "bg-gray-300 shadow-md flex flex-col justify-center items-center rounded-md w-full px-2 ">
       <div class = "w-full mb-3 mt-3 flex flex-row">
         <div class = "rounded-md text-white bg-black w-60 flex flex-col justify-center items-center lg:text-2xl text-xl">
           <div class = "ml-2 w-full h-9">Name</div> 
           <div class = "ml-2 w-full h-9">Color</div> 
           <div class = "ml-2 w-full h-9">Size</div> 
           <div class = "ml-2 w-full h-9">Material</div> 
           <div class = "ml-2 w-full h-9">Made In</div> 
           <div class = "ml-2 w-full h-9">Amount</div> 
         </div>
         <div class = "">
           <div class = "mx-3 h-9 flex items-center">
             <div class = "text-2xl font-medium pb-2">{{Product.name}}</div>
           </div>
           <div class = "mx-3 h-9 flex items-center pb-1">
              <button class = "w-8 h-8 border-black border-2 rounded-md bg-white focus:bg-gray-400 focus:outline-none" v-for="(color, idx) in Colors" :key="idx" v-bind:style="{ backgroundColor: color}"></button> 
           </div>
           <div class = "mx-3 h-9 flex items-center pt-1">
             <button class = "w-8 h-8 border-black border-2 rounded-md focus:bg-black focus:text-white hover:bg-black hover:text-white" v-for="(size, index) in Sizes" :key="index">{{size}}</button>
            </div>
           <div class = "x-3 h-9 flex items-center">
             <div class = "text-2xl font-medium">{{Product.fabric}}</div>
           </div>
           <div class = "mx-3 h-9 flex items-center">
             <div class = "text-2xl font-medium">{{Product.manufacturer}}</div>
           </div>
           <div class = "mx-3 h-9 flex items-center pt-1">
             <div class = "w-auto mr-2 flex items-center justify-center">
              <button @click.prevent="decrenemtAmount()" class="w-5 h-5 outline-none rounded-3xl flex items-center justify-center text-white bg-gray-700 text-md hover:bg-gray-800 focus:outline-none">-</button>
             </div>
             <input class = "w-20 bg-white h-9 rounded-md flex justify-center items-center text-xl overflow-x-auto" v-model="Amount" />
             <div class = "w-auto ml-2">
              <button @click.prevent="incremectAmount()" class="w-5 h-5 outline-none rounded-3xl flex items-center justify-center text-white bg-gray-700 text-md hover:bg-gray-800 focus:outline-none">+</button>
             </div>
           </div>
         </div>
         <div class = "w-full py-4 flex justify-center items-center text-white lg:px-20 px-10">
           <div class = "h-full bg-black flex flex-col rounded-xl justify-start items-center px-4 py-2">
             <div class = "lg:text-2xl text-md font-bold sm:mb-0 mb-4">About product</div>
             <div class = " h-full lg:text-2xl text-md flex items-center pb-8">{{Product.description}}</div>
           </div>
         </div>
       </div>
        <div class = "w-full mb-3 lg:pl-20 pl-3 flex flex-row rounded-md bg-black py-2"> 
         <div class = " text-white flex items-center justify-center flex-col"> 
           <div class = "text-2xl tracking-widest">USA</div> 
           <div class = "text-2xl tracking-widest">China</div> 
           <div class = "text-2xl tracking-widest">Europa</div> 
         </div> 
        <span class = "lg:ml-20 lg:mr-20 ml-6 mr-6 h-100 bg-white w-1 mb-2 mt-2"></span> 
        <div class = "w-full flex flex-col"> 
          <div class="grid grid-cols-10 gap-full gap-y-2 text-white">
            <div>2</div>
            <div>4</div>
            <div>6</div>
            <div>8</div>
            <div>10</div>
            <div>12</div>
            <div>14</div>
            <div>16</div>
            <div>18</div>
            <div>20</div>
            <div>5</div>
            <div>7</div>
            <div>9</div>
            <div>11</div>
            <div>13</div>
            <div>15</div>
            <div>17</div>
            <div>19</div>
            <div>21</div>
            <div>23</div>
            <div>2</div>
            <div>4</div>
            <div>6</div>
            <div>8</div>
            <div>10</div>
            <div>12</div>
            <div>14</div>
            <div>16</div>
            <div>18</div>
            <div>20</div>
          </div> 
        </div> 
       </div> 
       <div class = "w-full mb-3 flex flex-row text-white">
         <div class = "w-full lg:mr-60 mr-20">
           <button @click.prevent="goBack()" class = "uppercase flex justify-center items-center w-full bg-black rounded-xl hover:bg-gray-900 py-2 focus:outline-none">Cancel</button>
         </div>
         <div class = "w-full">
           <button @click.prevent="openCart()" class = "uppercase flex justify-center items-center w-full bg-black rounded-xl hover:bg-gray-900 py-2 focus:outline-none ">put in Cart </button>
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
        for(let i = 0; i < this.Product.sub_products.length - 1; i++) temp[i] = this.Product.sub_products[i].size
        return unique(temp)
      },
      Colors: function (){
        let temp = []
        for(let i = 0; i < this.Product.sub_products.length - 1; i++) temp[i] = this.Product.sub_products[i].color
        return unique(temp)
      },
      Images: function (){
        return this.Product.images
      }
    },
    mounted() { },
    methods:{
      goBack(){
        this.$router.push("/");
      },
      incremectAmount(){
        if(this.Product.sub_products.length > this.Amount)
          this.Amount++
      },
      decrenemtAmount(){
        if(0 < this.Amount)
          this.Amount--
      },
      openCart(){
        this.$router.push('/cart/'+ this.Product.id)
      },
      SelectImage(index){
        MainImage.src = `http://localhost:8080/static/${this.Product.images[index].url}`
      },


      SelectColor(ClothesColor){
        
      },
      SelectSize(ClotheSize){

      }
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