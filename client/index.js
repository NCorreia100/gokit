import React from 'react';
import ReactDOM from 'react-dom';


class App extends React.Component{
    constructor(props){
        super(props);
        this.state={}
    }
    render(){
        return(<h1>Hello! Hello!</h1>)
    }
}


ReactDOM.render(document.getElementById('gokit-app'),<App/>);

