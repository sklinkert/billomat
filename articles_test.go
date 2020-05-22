package billomat

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToArticles(t *testing.T) {
	expected := `<articles type="array" page="1" per_page="100" total="1"><article><id type="integer">812057</id><created type="datetime">2020-05-13T01:35:04+02:00</created><updated type="datetime">2020-05-16T12:27:35+02:00</updated><archived type="bool">0</archived><unit_id type="integer"/><article_number>D-1830-54</article_number><number type="integer">54</number><number_pre>D-1830-</number_pre><number_length type="integer">0</number_length><type/><title>8/6/8 Doppelstabmatte 1830 mm feuerverzinkt und RAL 7016 (Anthrazit) nach DIN EN 1461</title><description>8/6/8 Doppelstabmatte 1830 mm feuerverzinkt und RAL 7016 (Anthrazit) nach DIN EN 1461</description><sales_price type="float">56.87</sales_price><sales_price2 type="float"/><sales_price3 type="float"/><sales_price4 type="float"/><sales_price5 type="float"/><currency_code>EUR</currency_code><tax_id type="integer">83653</tax_id><revenue_account_number type="integer"/><cost_center type="integer"/><purchase_price type="float"/><purchase_price_net_gross>NET</purchase_price_net_gross><supplier_id type="integer"/><customfield/><article-property-values type="array"/></article></articles>`
	articles, err := toArticles([]byte(expected[:]))
	assert.NoError(t, err)
	assert.True(t.Fatalf, articles != nil)
	assert.EqualInt(t.Fatalf, 1, len(*articles))
	assert.EqualInt(t, 812057, (*articles)[0].ID)
}
