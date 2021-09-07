<?php
require 'vendor/autoload.php';
use PHPUnit\Framework\TestCase;

class SearchInsertTest extends TestCase
{

    public function testSearchInsert()
    {
        $Solution = new Practice\Solution();
		$testData = array(
			array(
				'nums' => array(1,3,5,6),
				'target' => 5,
				'expected' => 2
			),
			array(
				'nums' => array(1,3,5,6),
				'target' => 2,
				'expected' => 1
			),
			array(
				'nums' => array(1,3,5,6),
				'target' => 7,
				'expected' => 4
			),
			array(
				'nums' => array(1,3,5,6),
				'target' => 0,
				'expected' => 0
			),
			array(
				'nums' => array(1),
				'target' => 0,
				'expected' => 0
			),
		);
		foreach ($testData as $values) {
			$this->assertEquals($Solution->SearchInsert($values['nums'], $values['target']), $values['expected']);
		}
    }
}
?>