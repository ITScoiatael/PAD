<template>
    
 <div class = "w-full flex justify-center items-center flex-row px-10 py-20 ">
   <div class = "w-full flex justify-center items-center flex-row ">
     <div class = "border-2 border-black w-300 h-330">
       <img src="" alt="photo" class = "w-full h-full">
     </div>
     <div class = "flex flex-col mx-3 justify-between items-center mr-40">

       <div class = "mb-2 w-20 h-20 shadow-md bg-gray-300">
         <img src="" alt= "цвет товара">
       </div>
       
       <div  class = "mb-2 w-20 h-20 shadow-md bg-gray-300">
         <img src="" alt= "цвет товара">
       </div>
       
       <div class = " mb-2 w-20 h-20 shadow-md bg-gray-300">
         <img src="" alt= "цвет товара">
       </div>
      
       <div class = " w-20 h-20 shadow-md bg-gray-300">
         <img src="" alt= "цвет товара">
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
             <div class = "text-2xl font-medium pb-2">3.8$</div>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center pb-1">
             <button class = "w-8 h-8 border-black border-2 rounded-md bg-white focus:bg-gray-400 focus:outline-none"></button>
             <button class = "w-8 h-8 border-black border-2 rounded-md ml-2 mr-2 bg-black focus:bg-gray-400 focus:outline-none"></button>
             <button class = "w-8 h-8 border-black border-2 rounded-md mr-2 bg-yellow-200 focus:bg-gray-400 focus:outline-none"></button>
             <button class = "w-8 h-8 border-black border-2 rounded-md bg-red-600 focus:bg-gray-400 focus:outline-none"></button>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center pt-1">
             <button class = "w-8 h-8 border-black border-2 rounded-md focus:bg-black focus:text-white hover:bg-black hover:text-white">41</button>
             <button class = "w-8 h-8 border-black border-2 rounded-md ml-2 mr-2 focus:bg-black focus:text-white hover:bg-black hover:text-white">42</button>
             <button class = "w-8 h-8 border-black border-2 rounded-md mr-2 focus:bg-black focus:text-white hover:bg-black hover:text-white">43</button>
             <button class = "w-8 h-8 border-black border-2 rounded-md focus:bg-black focus:text-white hover:bg-black hover:text-white">44</button>
            </div>
           
           <div class = "w-full ml-3 h-9 flex items-center">
             <div class = "text-2xl font-medium">Cotton</div>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center">
             <div class = "text-2xl font-medium">China</div>
           </div>
           <div class = "w-full ml-3 h-9 flex items-center pt-1">
             <div class = "w-auto mr-2 flex items-center justify-center">
              <button class="w-5 h-5 outline-none rounded-3xl flex items-center justify-center text-white bg-gray-700 text-md hover:bg-gray-800 focus:outline-none">-</button>
             </div>
             <div class = "w-20 bg-white h-9 rounded-md flex justify-center items-center text-xl overflow-x-auto"></div>
             <div class = "w-auto ml-2">
              <button class="w-5 h-5 outline-none rounded-3xl flex items-center justify-center text-white bg-gray-700 text-md hover:bg-gray-800 focus:outline-none">+</button>
             </div>
           </div>
         </div>
         <div class = "w-full py-4 flex justify-center items-center text-white ">
           <div class = "w-300 h-full bg-black flex flex-col rounded-xl justify-start items-center px-4 py-2">
             <div class = "w-full text-2xl">About product</div>
             <div class = " w-full h-full text-xl flex items-center pb-8">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum ut ex at lacus bibendum sollicitudin. </div>
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
           <button class = "uppercase flex justify-center items-center w-full bg-black rounded-xl hover:bg-gray-900 py-2 focus:outline-none">Cancel</button>
         </div>
         <div class = "w-full">
           <button class = "uppercase flex justify-center items-center w-full bg-black rounded-xl hover:bg-gray-900 py-2 focus:outline-none ">buy</button>
         </div>
       </div>
     </div>
   </div>
 </div>
</template>