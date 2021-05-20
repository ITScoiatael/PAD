<template>
    <section>
        <HeroImage/>
        <Categories/>
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