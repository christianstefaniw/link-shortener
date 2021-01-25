import './App.css';
import {Component} from 'react'
import {Container} from "reactstrap";

import ShortenUrlForm from "./components/shorten_url/shorten_url_form";
import ShortenedUrls from "./components/shortened_urls/shortened_urls";

class App extends Component {

    constructor(props) {
        super(props);
        this.linksWereShortened = this.linksWereShortened.bind(this)
        this.state = {
            links: undefined

        }
    }

    linksWereShortened(links) {
        this.setState({
            links: links
        })
    }

    render() {
        return (
            <div>
                <Container>
                    <ShortenUrlForm onLinkChange={this.linksWereShortened}/>
                    <ShortenedUrls links={this.state.links}/>
                </Container>
            </div>
        );
    }
}

export default App;
