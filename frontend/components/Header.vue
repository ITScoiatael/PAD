<template>
    <nav class="bg-gray-400 py-2 md:py-4">
        <div class="container px-4 mx-auto md:flex md:items-center">
            <OurLogo/>
            <div class="hidden md:flex flex-col md:flex-row md:ml-auto mt-3 md:mt-0" id="navbar-collapse">
                <div class="content-around ">
                    <NuxtLink  to="./"
                        class="p-2 lg:px-4 md:mx-2 text-black-900 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300">Home</NuxtLink >
                    <a href="https://github.com/ITScoiatael/PAD" target="_blank"
                        class="p-2 lg:px-4 md:mx-2 text-black-900 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300">About
                        us</a>
                    <NuxtLink  to="Catalog"
                        class="p-2 lg:px-4 md:mx-2 text-black-900 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300">Catalog</NuxtLink>
                </div>
            </div>
        </div>
    </nav>
</template>