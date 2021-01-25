import {Component} from 'react'
import {Card, Container} from "reactstrap";
import ShortenedUrl from "./shortened_url";

class ShortenedUrls extends Component{
    render() {
        return (
            <Card id="shortened-urls">
                <Container>
                    {this.props.links !== undefined ? this.props.links.map((link, i) => <ShortenedUrl backgroundColor={i % 2 === 0 ? 'white' : '#F8F8F8'} link={link} key={i}/>) : null}
                </Container>
            </Card>
        );
    }
}

export default ShortenedUrls