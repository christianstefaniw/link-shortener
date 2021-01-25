import {Component} from 'react'
import {Button} from "reactstrap";

class ShortenUrlButton extends Component{

    constructor(props) {
        super(props);
        this.props.onSubmit.bind(this)
    }


    render() {
        return (
            <Button onClick={this.props.onSubmit} color="primary" size="lg" block>Shorten</Button>
        );
    }
}

export default ShortenUrlButton