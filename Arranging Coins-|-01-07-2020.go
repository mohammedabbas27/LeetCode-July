func arrangeCoins(n int) int {
    if(n <=1){
        return n
    }
    
    counter := 0
    store := n
    i:=1
    for i=1;i<n;i++{
        if(store < i){
           break
        }
        store -= i
        counter+=(n-(n-i))
    }
    return i-1 
}
