<?php

namespace Practice;

class Solution {
    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function SearchInsert($nums, $target) {
        if ($target < $nums[0]) {
            return 0;
        }
        $last = $nums[0];
        foreach ($nums as $key => $value) {
            if($target <= $value && $target >= $last){
                return $key;
            }
            $last = $value;
        }
        return count($nums);
    }
}

?>