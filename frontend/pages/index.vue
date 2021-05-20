<template>
    <section class="bg-GrayImage">
        <HeroImage/>
        <Categories/>
        <LineWithLogo/>
        <Services />
    </section>
</template>

<script>
import config from '../env'

export default {
    mounted(){
        this.env = config;
        console.log(this.env)
    }
}
</script>