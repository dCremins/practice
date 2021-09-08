---
to: <%= h.changeCase.pascal(name) %>/<%= h.changeCase.pascal(name) %>Test.php
---
<?php
require 'vendor/autoload.php';
use PHPUnit\Framework\TestCase;

class <%= h.changeCase.pascal(name) %>Test extends TestCase
{
    protected $stack;

    protected function setUp(): void
    {
        $this->Solution = new Practice\Solution();
    }
	
    public function test<%= h.changeCase.pascal(name) %>()
    {
		$testData = array(
			[0]=>array(
				'input' => 'input',
				'expected' => 'expected'
			)
		);
		foreach ($testData as $values) {
			$this->assertEquals($this->Solution-><%= h.changeCase.pascal(name) %>($values['input']), $values['expected']);
		}
    }
}
?>