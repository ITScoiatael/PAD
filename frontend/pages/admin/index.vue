<template>
    <section class="">
        <div class = "flex w-full flex-col h-full h-screen">
            <div class = "w-auto flex justify-center items-center p-10 flex-row">
                <span class = "bg-black h-1 w-full mr-4"></span>
                <div class = "w-36 "><Logo/> </div>
                <span class = "bg-black h-1 w-full mr-4"></span>
            </div>
        </div>
        <div class = "w-full mt-20">
            <div class = "flex flex-col w-full px-10">
                <div class = "text-black uppercase font-bold">Count</div>
                <button class = "w-10 h-10 border-2 border-black mt-2 rounded-md items-center justify-center">+</button>
            </div>
            <div class = "flex flex-col w-full px-10 mt-5">
                <div class = "text-black uppercase font-bold">Size</div>
                <button class = "w-10 h-10 border-2 border-black mt-2 rounded-md items-center justify-center">+</button>
            </div>
            <div class = "flex flex-col w-full px-10 mt-5">
                <div class = "text-black uppercase font-bold">Color</div>
                <button class = "w-10 h-10 border-2 border-black mt-2 rounded-md items-center justify-center">+</button>
            </div>
        </div>
        <div class = "w-full mb-10 mt-10 flex flex-col items-center justify-center px-60">
            <div class = "w-full flex justify-center"><input type="text" name="" id="" class = "border-2 border-black rounded-md w-full py-2 px-2 pb-10" placeholder="Information"></div>
            <div class = "mt-5 ">
                <button class = "px-6 py-1 bg-gray-300 hover:bg-gray-400 w-full border-2 border-black uppercase font-bold">Enter</button>
            </div>
        </div>
    </section>
</template>