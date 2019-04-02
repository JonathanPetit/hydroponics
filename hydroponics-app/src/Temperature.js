import React, { Component } from 'react';
import axios from 'axios'

class Temperature extends Component {
    constructor(props) {
        super(props);
        this.state = {
          temperatures: []
        };
    }

    componentDidMount() {
        axios.get("http://localhost:8080/temperatures")
        .then(res => {
            const temps = res.data;
            this.setState({temperatures: temps});})
        .catch(error => console.log("Error Temp"));
    }

    render() {
        return (
            <div>{this.state.temperatures.map(temp => <div>{Date(temp.savedAt)}</div>)}</div>
        )
    }
}

export default Temperature;