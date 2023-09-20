package org.example;

import java.util.*;

public class Main {
    public static void main(String[] args) {

        List<Integer> list1 = new ArrayList<>();

        Set<Integer> set1 = new TreeSet<>();

        list1.add(23);
        list1.add(25);
        list1.add(25);
        list1.add(25);
        list1.add(25);

        set1.add(23);
        set1.add(21);
        set1.add(223);
        set1.add(2453);
        set1.add(232);
        set1.add(231);

        for(Integer i : set1) {
            System.out.println(i);
        }


    }
}