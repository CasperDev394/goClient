<template>
    <div class="root">
        <input type="button" @click="methodsGetPing()" value="">
        <input type="text" :value="ping">
        <ul>
            <li 
                v-for="(value, key, index) in ping"
                :key="index"    
            >
                {{ `${key} ${value}`}}
            </li>
        </ul>
        <br>
        <input type="button" @click="methodsGetSimpleSvsC()" value="">
        <input type="text" :value="SvsC">
        <ul>
            <li 
                v-for="(value, key, index) in SvsC"
                :key="index"    
            >
                {{ `${key} ${value}`}}
            </li>
        </ul>
        <br>
        <input type="button" @click="TokenPrice()" value="">
        <input type="text" :value="TokenP">
        <ul>
            <li 
                v-for="(value, key, index) in TokenP"
                :key="index"    
            >
                {{ `${key} ${value}`}}
            </li>
        </ul>
        <br>
        <input type="button" @click="CoinsList()" value="">
        <input type="text" :value="CoinList">
        <ul>
            <li 
                v-for="(value, key, index) in CoinList"
                :key="index"    
            >
                {{ `${key} ${value}`}}
            </li>
        </ul>
        <br>
        <input type="button" @click="CoinHistory()" value="">
        <input type="text" :value="History">
        <ul>
            <li 
                v-for="(value, key, index) in History"
                :key="index"    
            >
                {{ `${key} ${value}`}}
            </li>
        </ul>
        <br>
        <input type="button" @click="CoinOhlc()" value="">
        <input type="text" :value="ohlc">
        <ul>
            <li 
                v-for="(value, key, index) in ohlc"
                :key="index"    
            >
                {{ `${key} ${value}`}}
            </li>
        </ul>
    </div>
</template>

<script>
    export default{
        name:'getPing',
        data(){
            return{
                ping:[],
                SvsC:[],
                TokenP:[],
                CoinList:[],
                History:[],
                ohlc:[]
            }
        },
        methods:{
            methodsGetPing(){
                console.log('Starting...') 
                fetch('/bot/ping', {
                    method: 'GET',
                })
                .then((resp)=>{
                    return resp.json()
                })
                .then((data)=>{
                    this.ping = data
                })
            },
            methodsGetSimpleSvsC(){
                console.log('Starting SvsC...') 
                fetch('/bot/simple/supported_vs_currencies', {
                    method: 'GET',
                })
                .then((resp)=>{
                    return resp.json()
                })
                .then((data)=>{
                    this.SvsC = data
                })
            },
            TokenPrice(){
                console.log('Starting TokenP...') 
                fetch('/bot/simple/token_price/1223/', {
                    method: 'GET',
                })
                .then((resp)=>{
                    return resp.json()
                })
                .then((data)=>{
                    this.TokenP = data
                })
            },
            CoinsList(){
                console.log('Starting TokenP...') 
                fetch('/bot/coins/list', {
                    method: 'GET',
                })
                .then((resp)=>{
                    return resp.json()
                })
                .then((data)=>{
                    this.CoinList = data
                })
            },
            CoinHistory(){
                console.log('Starting TokenP...') 
                fetch('/bot/coins/bitcoin/history?date=18-10-2022', {
                    method: 'GET',
                })
                .then((resp)=>{
                    return resp.json()
                })
                .then((data)=>{
                    this.History = data
                })
            },
            CoinOhlc(){
                console.log('Starting TokenP...') 
                fetch('/bot/coins/bitcoin/ohlc?days=1&vs_currency=usd', {
                    method: 'GET',
                })
                .then((resp)=>{
                    return resp.json()
                })
                .then((data)=>{
                    this.ohlc = data
                })
            }
            
        }
    }
</script>
