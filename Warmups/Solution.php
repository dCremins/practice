<?php

namespace Practice;

class Solution {
    /**
     * @param Integer $n
     * @param Integer[] $ar[n]
     * @return Integer
     */
    function sockMerchant($n, $ar) {
        $countArray = array_count_values($ar);
        $pairs = 0;
        foreach($countArray as $key=>$value) {
            $keyPairs = intdiv($value, 2);
            $pairs += $keyPairs;
        }
        return $pairs;
    }

    /**
     * @param Integer $steps
     * @param String $path
     * @return Integer
     */
    function countingValleys($steps, $path) {
        $pattern = "/D+|U+/";
        $level = 0;
        $valleys = 0;
        $inValley = false;
        if(preg_match_all($pattern,$path,$matches)) {
            foreach($matches[0] as $change){
                $difference = strlen($change);
                if($change[0] == 'D'){
                    $level -= $difference;
                    if($level <0){
                        $inValley = true;
                    }
                } else {
                    $level += $difference;
                }
                if($inValley && $level >= 0){
                    $valleys += 1;
                    $inValley = false;
                }
            }
        }
        return $valleys;
    }
}

?>