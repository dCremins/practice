---
to: <%= h.changeCase.pascal(name) %>/test.php
---
<?php
use PHPUnit\Framework\TestCase;
use index

class <%= h.changeCase.pascal(name) %>Test extends TestCase
{
    public function setUp()
    {
        $this->Solution = new Solution();
    }

    public function test<%= h.changeCase.pascal(name) %>()
    {
		$testData = array(
			[0]=>array(
				'input' => 'input',
				'expected' => 'expected'
			)
		)
		foreach ($testData as $values) {
			$this->assertEquals($this->Solution-><%= h.changeCase.pascal(name) %>($values['input']), $values['expected']);
		}
    }
}
?>