<?php
require 'vendor/autoload.php';
use PHPUnit\Framework\TestCase;

class WarmupTest extends TestCase
{
    protected $stack;

    protected function setUp(): void
    {
        $this->Solution = new Practice\Solution();
    }
    public function testOne()
    {
		$testData = array(
			array(
				'n' => 9,
				'ar' => array(10, 20, 20, 10, 10, 30, 50, 10, 20),
				'expected' => 3
			),
			array(
				'n' => 7,
				'ar' => array(1,2,1,2,1,3,2),
				'expected' => 2
			),
			array(
				'n' => 10,
				'ar' => array(1,1,3,1,2,1,3,3,3,3),
				'expected' => 4
			)
		);
		foreach ($testData as $values) {
			$this->assertEquals($this->Solution->sockMerchant($values['n'], $values['ar']), $values['expected']);
		}
    }
    public function testTwo()
    {
		$testData = array(
			array(
				'steps' => 8,
				'path' => 'UDDDUDUU',
				'expected' => 1
			),
			array(
				'steps' => 12,
				'path' => 'DDUUDDUDUUUD',
				'expected' => 2
			),
		);
		foreach ($testData as $values) {
			$this->assertEquals($this->Solution->countingValleys($values['steps'], $values['path']), $values['expected']);
		}
    }
}
?>